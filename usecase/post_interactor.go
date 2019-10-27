package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
	"github.com/naoki85/my-blog-api-sam/util"
	"io"
	"log"
	"sort"
	"time"
)

type PostInteractor struct {
	PostRepository      PostRepository
	IdCounterRepository IdCounterRepository
}

type PostInteractorCreateParams struct {
	Id          int    `json:"id"`
	Category    string `json:"category"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	ImageUrl    string `json:"imageUrl"`
	Active      string `json:"active"`
	PublishedAt string `json:"publishedAt"`
	ImageBody   io.Reader
}

func (interactor *PostInteractor) Index(page int, all bool) (posts model.Posts, count int, err error) {
	posts, count, err = interactor.PostRepository.All()
	if err != nil {
		log.Printf("%s", err.Error())
	}
	var retPosts model.Posts
	layout := "2006-01-02 15:04:05"
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Tokyo")
	for _, post := range posts {
		if !all {
			t, err := time.ParseInLocation(layout, post.PublishedAt, loc)
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

func (interactor *PostInteractor) Create(params PostInteractorCreateParams) (err error) {
	maxId, err := interactor.IdCounterRepository.FindMaxIdByIdentifier("Posts")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	newId := maxId + 1
	_, err = interactor.IdCounterRepository.UpdateMaxIdByIdentifier("Posts", newId)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	var inputParams = repository.PostCreateParams{
		Id:          newId,
		UserId:      1,
		Category:    params.Category,
		Title:       params.Title,
		Content:     params.Content,
		ImageUrl:    params.ImageUrl,
		Active:      params.Active,
		PublishedAt: params.PublishedAt,
	}
	err = interactor.PostRepository.Create(inputParams)
	return
}

func (interactor *PostInteractor) Update(params PostInteractorCreateParams) (err error) {
	var inputParams = repository.PostCreateParams{
		Id:          params.Id,
		UserId:      1,
		Category:    params.Category,
		Title:       params.Title,
		Content:     params.Content,
		ImageUrl:    params.ImageUrl,
		Active:      params.Active,
		PublishedAt: params.PublishedAt,
	}
	err = interactor.PostRepository.Update(inputParams)
	return
}

func (interactor *PostInteractor) Delete(id int) (err error) {
	err = interactor.PostRepository.Delete(id)
	return
}
