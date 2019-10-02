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

	recommendedBooks, err := interactor.RecommendedBookRepository.All()
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if len(recommendedBooks) != 5 {
		t.Fatalf("Fail expected: 5, got: %v", len(recommendedBooks))
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
