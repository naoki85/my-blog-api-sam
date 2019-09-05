package controller

import (
	"encoding/json"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/interface/database"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"strings"
)

type PostController struct {
	Interactor usecase.PostInteractor
}

func NewPostController(sqlHandler database.SqlHandler) *PostController {
	return &PostController{
		Interactor: usecase.PostInteractor{
			PostRepository: &database.PostRepository{
				SqlHandler: sqlHandler,
			},
			PostCategoryRepository: &database.PostCategoryRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *PostController) Index(page int) ([]byte, int) {
	posts, err := controller.Interactor.Index(page)
	if err != nil {
		return []byte{}, config.NotFoundStatus
	}

	var retPosts model.Posts

	if len(posts) == 0 {
		retPosts = model.Posts{}
	}

	for _, p := range posts {
		if p.ImageUrl == "" {
			p.ImageUrl = "https://s3-ap-northeast-1.amazonaws.com/bookrecorder-image/commons/default_user_icon.png"
		} else {
			p.ImageUrl = "http://d29xhtkvbwm2ne.cloudfront.net/" + p.ImageUrl
		}
		p.PublishedAt = strings.Split(p.PublishedAt, "T")[0]

		retPosts = append(retPosts, p)
	}

	count, err := controller.Interactor.GetPostsCount()
	if err != nil {
		return []byte{}, config.InternalServerErrorStatus
	}

	totalPage := count / 10
	mod := count % 10
	if mod != 0 {
		totalPage = totalPage + 1
	}

	data := struct {
		TotalPage int
		Posts     model.Posts
	}{totalPage, retPosts}
	resp, err := json.Marshal(data)
	if err != nil {
		return resp, config.InternalServerErrorStatus
	}

	return resp, config.SuccessStatus
}

func (controller *PostController) Show(id int) ([]byte, int) {
	post, err := controller.Interactor.FindById(id)
	if err != nil || post.Id == 0 {
		return []byte{}, config.NotFoundStatus
	}
	resp, err := json.Marshal(post)
	if err != nil {
		return resp, config.InternalServerErrorStatus
	}

	return resp, config.SuccessStatus
}
