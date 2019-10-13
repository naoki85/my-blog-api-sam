package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/util"
	"log"
	"sort"
	"time"
)

type PostInteractor struct {
	PostRepository      PostRepository
	IdCounterRepository IdCounterRepository
}

func (interactor *PostInteractor) Index(page int, all bool) (posts model.Posts, count int, err error) {
	posts, count, err = interactor.PostRepository.All()
	if err != nil {
		log.Printf("%s", err.Error())
	}
	var retPosts model.Posts
	layout := "2006-01-02 15:04:05"
	now := time.Now()
	for _, post := range posts {
		if !all {
			t, err := time.Parse(layout, post.PublishedAt)
			if err != nil {
				log.Printf("%s", err.Error())
				continue
			}
			if t.After(now) {
				continue
			}
		}
		if post.ImageUrl == "-" {
			post.ImageUrl = "https://s3-ap-northeast-1.amazonaws.com/bookrecorder-image/commons/default_user_icon.png"
		} else {
			post.ImageUrl = "http://d29xhtkvbwm2ne.cloudfront.net/" + post.ImageUrl
		}
		retPosts = append(retPosts, post)
	}
	count = len(retPosts) / 10
	if len(retPosts)%10 != 0 {
		count++
	}

	sort.Slice(retPosts, func(i, j int) bool { return retPosts[i].Id > retPosts[j].Id })
	start := util.CompareInt("max", (page-1)*10, 0)
	end := util.CompareInt("min", page*10, len(retPosts))

	var ret []model.Post
	ret = retPosts[start:end]

	return ret, count, nil
}

func (interactor *PostInteractor) FindById(id int) (post model.Post, err error) {
	post, err = interactor.PostRepository.FindById(id)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}
	post.ImageUrl = "http://d29xhtkvbwm2ne.cloudfront.net/" + post.ImageUrl

	return
}
