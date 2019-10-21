package controller

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"log"
)

type PostController struct {
	Interactor usecase.PostInteractor
}

func NewPostController(dynamoDbHandler *dynamodb.DynamoDB) *PostController {
	return &PostController{
		Interactor: usecase.PostInteractor{
			PostRepository: &repository.PostRepository{
				DynamoDBHandler: dynamoDbHandler,
			},
			IdCounterRepository: &repository.IdCounterRepository{
				DynamoDBHandler: dynamoDbHandler,
			},
		},
	}
}

func (controller *PostController) Index(page int, all bool) ([]byte, int) {
	posts, count, err := controller.Interactor.Index(page, all)
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.NotFoundStatus
	}

	data := struct {
		TotalPage int
		Posts     model.Posts
	}{count, posts}
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}

	return resp, config.SuccessStatus
}

func (controller *PostController) Show(id int, format string) ([]byte, int) {
	post, err := controller.Interactor.FindById(id)
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.NotFoundStatus
	}
	if post.Id == 0 {
		log.Println("post is not found")
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

func (controller *PostController) Create(params usecase.PostInteractorCreateParams) ([]byte, int) {
	err := controller.Interactor.Create(params)
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.InvalidParameterStatus
	}

	data := struct {
		Message string
	}{"success"}
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}
	return resp, config.SuccessStatus
}

func (controller *PostController) Update(params usecase.PostInteractorCreateParams) ([]byte, int) {
	err := controller.Interactor.Update(params)
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.InvalidParameterStatus
	}

	data := struct {
		Message string
	}{"success"}
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}
	return resp, config.SuccessStatus
}

func ogpHtml(post model.Post) ([]byte, error) {
	var html string
	var content string
	if len(post.Content) < 160 {
		content = post.Content
	} else {
		content = post.Content[:160]
	}
	html = fmt.Sprint(`<!DOCTYPE html><html lang="ja"><head><title>naoki85 のブログ</title>`)
	html = html + fmt.Sprint(`<link rel="shortcut icon" type="image/x-icon" href="//d1mtswcgj7q8jb.cloudfront.net/assets/commons/favicon-de7e93026202b78eff192fcda074c780a0f6cfb4e11f591eab829a0ef91c965f.ico" /><meta charset="utf-8">`)
	html = html + fmt.Sprintf(`<title>%s</title>`, post.Title)
	html = html + fmt.Sprint(`<meta name="keywords" content="book,本">`)
	html = html + fmt.Sprintf(`<meta property="og:title" content="%s">`, post.Title)
	html = html + fmt.Sprint(`<meta property="og:type" content="website">`)
	html = html + fmt.Sprintf(`<meta property="og:url" content="https://blog.naoki85.me/posts/%d">`, post.Id)
	html = html + fmt.Sprintf(`<meta property="og:image" content="%s">`, post.ImageUrl)
	html = html + fmt.Sprint(`<meta property="og:site_name" content="naoki85 のブログ">`)
	html = html + fmt.Sprintf(`<meta property="og:description" content="%s">`, content)
	html = html + fmt.Sprint(`<meta property="og:locale" content="ja_JP">`)
	html = html + fmt.Sprint(`<meta name="twitter:card" content="summary">`)
	html = html + fmt.Sprint(`<meta name="twitter:site" content="@tony_201612">`)
	html = html + fmt.Sprint(`<meta name="twitter:creator" content="@tony_201612"></head>`)
	html = html + fmt.Sprintf(`<body><script>window.location = "https://blog.naoki85.me/posts/%d";</script></body></html>`, post.Id)
	return []byte(html), nil
}
