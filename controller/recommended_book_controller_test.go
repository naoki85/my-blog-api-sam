package controller

import (
	"encoding/json"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"testing"
)

func TestShouldFindAllRecommendedBooks(t *testing.T) {
	sqlHandler, tearDown := SetupTest()
	dynamoDbHandler, _ := testSupport.NewDynamoDbHandler()
	defer tearDown()
	controller := NewRecommendedBookController(sqlHandler, dynamoDbHandler)
	_, _ = sqlHandler.Execute("INSERT INTO `recommended_books` VALUES (1,6,'http://test.example.com/hoge','http://test.example.com/hoge.png','http://test.example.com/hoge.png',0,'2018-12-30 17:00:14','2018-12-30 17:00:25')")
	_, _ = sqlHandler.Execute("INSERT INTO `recommended_books` VALUES (2,6,'http://test.example.com/hoge','http://test.example.com/hoge.png','http://test.example.com/hoge.png',0,'2018-12-30 17:00:14','2018-12-30 17:00:25')")
	_, _ = sqlHandler.Execute("INSERT INTO `recommended_books` VALUES (3,6,'http://test.example.com/hoge','http://test.example.com/hoge.png','http://test.example.com/hoge.png',0,'2018-12-30 17:00:14','2018-12-30 17:00:25')")
	_, _ = sqlHandler.Execute("INSERT INTO `recommended_books` VALUES (4,6,'http://test.example.com/hoge','http://test.example.com/hoge.png','http://test.example.com/hoge.png',0,'2018-12-30 17:00:14','2018-12-30 17:00:25')")

	recommendedBooks, status := controller.Index()
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}
	var res struct {
		RecommendedBooks model.RecommendedBooks
	}
	err := json.Unmarshal(recommendedBooks, &res)
	if err != nil {
		t.Fatalf("Response could not pasred: %s", err.Error())
	}
	if len(res.RecommendedBooks) != 4 {
		t.Fatalf("Fail expected length: 4, got: %v", res)
	}
}

func TestShouldCreateRecommendedBook(t *testing.T) {
	sqlHandler, tearDown := SetupTest()
	dynamoDbHandler, _ := testSupport.NewDynamoDbHandler()
	defer tearDown()
	controller := NewRecommendedBookController(sqlHandler, dynamoDbHandler)

	params := usecase.RecommendedBookInteractorCreateParams{
		Link:      "http://test.example.com/hoge",
		ImageUrl:  "http://test.example.com/hoge.png",
		ButtonUrl: "http://test.example.com/hoge.png",
	}

	_, status := controller.Create(params)
	if status != config.SuccessStatus {
		t.Fatalf("Should get 200 status, but got: %d", status)
	}
}
