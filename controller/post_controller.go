package controller

import (
	"encoding/json"
	"fmt"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"log"
	"strings"
)

type PostController struct {
	Interactor usecase.PostInteractor
}

func NewPostController(sqlHandler repository.SqlHandler) *PostController {
	return &PostController{
		Interactor: usecase.PostInteractor{
			PostRepository: &repository.PostRepository{
				SqlHandler: sqlHandler,
			},
			PostCategoryRepository: &repository.PostCategoryRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *PostController) Index(page int) ([]byte, int) {
	posts, err := controller.Interactor.Index(page)
	if err != nil {
		log.Printf("%s", err.Error())
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
		log.Printf("%s", err.Error())
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
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}

	return resp, config.SuccessStatus
}

func (controller *PostController) Show(id int, format string) ([]byte, int) {
	post, err := controller.Interactor.FindById(id)
	if err != nil || post.Id == 0 {
		log.Printf("%s", err.Error())
		return []byte{}, config.NotFoundStatus
	}

	var resp []byte
	var err2 error
	if format == "json" {
		resp, err2 = json.Marshal(post)
	} else {
		resp, err2 = ogpHtml(post)
	}

	if err2 != nil {
		log.Printf("%s", err2.Error())
		return resp, config.InternalServerErrorStatus
	}

	return resp, config.SuccessStatus
}

func ogpHtml(post model.Post) ([]byte, error) {
	var html string
	html = fmt.Sprint(`<!DOCTYPE html><html lang="ja"><head><title>naoki85 のブログ</title>`)
	html = html + fmt.Sprint(`<link rel="shortcut icon" type="image/x-icon" href="//d1mtswcgj7q8jb.cloudfront.net/assets/commons/favicon-de7e93026202b78eff192fcda074c780a0f6cfb4e11f591eab829a0ef91c965f.ico" /><meta charset="utf-8">`)
	html = html + fmt.Sprintf(`<title>%s</title>`, post.Title)
	html = html + fmt.Sprint(`<meta name="keywords" content="book,本">`)
	html = html + fmt.Sprintf(`<meta property="og:title" content="%s">`, post.Title)
	html = html + fmt.Sprint(`<meta property="og:type" content="website">`)
	html = html + fmt.Sprintf(`<meta property="og:url" content="https://blog.naoki85.me/posts/%d">`, post.Id)
	html = html + fmt.Sprintf(`<meta property="og:image" content="%s">`, post.ImageUrl)
	html = html + fmt.Sprint(`<meta property="og:site_name" content="naoki85 のブログ">`)
	html = html + fmt.Sprintf(`<meta property="og:description" content="%s">`, post.Content[:160])
	html = html + fmt.Sprint(`<meta property="og:locale" content="ja_JP">`)
	html = html + fmt.Sprint(`<meta name="twitter:card" content="summary">`)
	html = html + fmt.Sprint(`<meta name="twitter:site" content="@tony_201612">`)
	html = html + fmt.Sprint(`<meta name="twitter:creator" content="@tony_201612"></head>`)
	html = html + fmt.Sprintf(`<body><script>window.location = "https://blog.naoki85.me/posts/%d";</script></body></html>`, post.Id)
	return []byte(html), nil
}
