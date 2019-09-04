package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"strings"
)

type PostInteractor struct {
	PostRepository         PostRepository
	PostCategoryRepository PostCategoryRepository
}

func (interactor *PostInteractor) Index(page int) (model.Posts, error) {
	posts, err := interactor.PostRepository.Index(page)
	var retPosts model.Posts
	for _, post := range posts {
		postCategory, err := interactor.PostCategoryRepository.FindById(post.PostCategoryId)
		if err != nil {
			continue
		}
		post.PostCategory = postCategory
		retPosts = append(retPosts, post)
	}
	return retPosts, err
}

func (interactor *PostInteractor) FindById(id int) (model.Post, error) {
	post, err := interactor.PostRepository.FindById(id)
	if err != nil {
		return post, err
	}
	post.PublishedAt = strings.Split(post.PublishedAt, "T")[0]
	post.ImageUrl = "http://d29xhtkvbwm2ne.cloudfront.net/" + post.ImageUrl
	return post, err
}

func (interactor *PostInteractor) GetPostsCount() (int, error) {
	count, err := interactor.PostRepository.GetPostsCount()
	if err != nil {
		return count, err
	}
	return count, err
}
