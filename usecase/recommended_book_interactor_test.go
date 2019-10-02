package usecase

import (
	"github.com/naoki85/my-blog-api-sam/repository"
	"github.com/naoki85/my-blog-api-sam/testSupport"
	"testing"
)

func TestShouldFindAllRecommendedBooks(t *testing.T) {
	sqlHandler, tearDown := SetupTest()
	dynamoDbHandler, _ := testSupport.NewDynamoDbHandler()
	defer tearDown()
	interactor := RecommendedBookInteractor{
		RecommendedBookRepository: &repository.RecommendedBookRepository{
			sqlHandler,
			dynamoDbHandler,
		},
	}

	_, _ = sqlHandler.Execute("INSERT INTO `recommended_books` VALUES (1,6,'http://test.example.com/hoge','http://test.example.com/hoge.png','http://test.example.com/hoge.png',0,'2018-12-30 17:00:14','2018-12-30 17:00:25')")
	_, _ = sqlHandler.Execute("INSERT INTO `recommended_books` VALUES (2,6,'http://test.example.com/hoge','http://test.example.com/hoge.png','http://test.example.com/hoge.png',0,'2018-12-30 17:00:14','2018-12-30 17:00:25')")
	_, _ = sqlHandler.Execute("INSERT INTO `recommended_books` VALUES (3,6,'http://test.example.com/hoge','http://test.example.com/hoge.png','http://test.example.com/hoge.png',0,'2018-12-30 17:00:14','2018-12-30 17:00:25')")
	_, _ = sqlHandler.Execute("INSERT INTO `recommended_books` VALUES (4,6,'http://test.example.com/hoge','http://test.example.com/hoge.png','http://test.example.com/hoge.png',0,'2018-12-30 17:00:14','2018-12-30 17:00:25')")

	recommendedBooks, err := interactor.RecommendedBookRepository.All(4)
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if len(recommendedBooks) != 4 {
		t.Fatalf("Fail expected: 4, got: %v", len(recommendedBooks))
	}
}

func TestShouldCreateRecommendedBook(t *testing.T) {
	sqlHandler, tearDown := SetupTest()
	dynamoDbHandler, _ := testSupport.NewDynamoDbHandler()
	defer tearDown()
	interactor := RecommendedBookInteractor{
		RecommendedBookRepository: &repository.RecommendedBookRepository{
			sqlHandler,
			dynamoDbHandler,
		},
	}

	params := RecommendedBookInteractorCreateParams{
		Link:      "http://test.example.com/hoge",
		ImageUrl:  "http://test.example.com/hoge.png",
		ButtonUrl: "http://test.example.com/hoge.png",
	}

	err := interactor.Create(params)
	if err != nil {
		t.Fatalf("Could not create recommended book: %s", err.Error())
	}
}
