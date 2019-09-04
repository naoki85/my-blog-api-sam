package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/naoki85/my-blog-api-sam/infrastructure"
	"github.com/naoki85/my-blog-api-sam/interface/database"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.Path == "/hello" {
		return hello(request)
	} else if request.Path == "/recommended_books" {
		return recommendedBooks(request)
	} else if request.Path == "/posts" {
		return posts(request)
	} else if regexp.MustCompile(`^/posts/(\d+)`).MatchString(request.Path) {
		return post(request)
	}
	return notFound()
}

func main() {
	lambda.Start(handler)
}

func hello(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp, err := http.Get(DefaultHTTPGetAddress)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if resp.StatusCode != 200 {
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if len(ip) == 0 {
		return events.APIGatewayProxyResponse{}, ErrNoIP
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %v", string(ip)),
		StatusCode: 200,
	}, nil
}

func recommendedBooks(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sqlHandler, _ := infrastructure.NewSqlHandler()
	limit := 4
	interactor := usecase.RecommendedBookInteractor{
		RecommendedBookRepository: &database.RecommendedBookRepository{
			SqlHandler: sqlHandler,
		},
	}
	recommendedBooks, err := interactor.All(limit)

	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return notFound()
	}

	data := struct {
		RecommendedBooks model.RecommendedBooks
	}{recommendedBooks}
	resp, err := json.Marshal(data)
	if err != nil {
		return internalServerError()
	}
	return events.APIGatewayProxyResponse{
		Body: fmt.Sprintf("%s", resp),
		StatusCode: 200,
	}, nil
}

func posts(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	page, err := strconv.Atoi(request.QueryStringParameters["page"])
	if err != nil {
		return invalidParameter()
	}

	sqlHandler, _ := infrastructure.NewSqlHandler()
	interactor := usecase.PostInteractor{
		PostRepository: &database.PostRepository{
			SqlHandler: sqlHandler,
		},
		PostCategoryRepository: &database.PostCategoryRepository{
			SqlHandler: sqlHandler,
		},
	}
	posts, err := interactor.Index(page)

	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return notFound()
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

	count, err := interactor.GetPostsCount()
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return internalServerError()
	}
	totalPage := count / 10
	mod := count % 10
	if mod != 0 {
		totalPage = totalPage + 1
	}

	data := struct {
		TotalPage int
		Posts model.Posts
	}{totalPage, posts}
	resp, err := json.Marshal(data)
	if err != nil {
		return internalServerError()
	}
	return events.APIGatewayProxyResponse{
		Body: fmt.Sprintf("%s", resp),
		StatusCode: 200,
	}, nil
}

func post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	postId, err := strconv.Atoi(request.PathParameters["id"])
	if err != nil {
		return invalidParameter()
	}

	sqlHandler, _ := infrastructure.NewSqlHandler()
	interactor := usecase.PostInteractor{
		PostRepository: &database.PostRepository{
			SqlHandler: sqlHandler,
		},
		PostCategoryRepository: &database.PostCategoryRepository{
			SqlHandler: sqlHandler,
		},
	}

	post, err := interactor.FindById(postId)
	if err != nil || post.Id == 0 {
		return notFound()
	}

	resp, err := json.Marshal(post)
	if err != nil {
		return internalServerError()
	}
	return events.APIGatewayProxyResponse{
		Body: fmt.Sprintf("%s", resp),
		StatusCode: 200,
	}, nil
}

func invalidParameter() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body: fmt.Sprint("Invalid Parameter"),
		StatusCode: 400,
	}, nil
}

func notFound() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body: fmt.Sprint("Not Found"),
		StatusCode: 404,
	}, nil
}

func internalServerError() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body: fmt.Sprint("Internal Server Error"),
		StatusCode: 500,
	}, nil
}
