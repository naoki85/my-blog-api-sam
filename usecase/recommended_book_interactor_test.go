package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"testing"
)

type MockRecommendedBookRepository struct {
}

func (repo *MockRecommendedBookRepository) All(int) (model.RecommendedBooks, error) {
	recommendedBooks := model.RecommendedBooks{
		model.RecommendedBook{Id: 1},
		model.RecommendedBook{Id: 2},
		model.RecommendedBook{Id: 3},
		model.RecommendedBook{Id: 4},
	}
	return recommendedBooks, nil
}

func TestShouldFindAllRecommendedBooks(t *testing.T) {
	repo := new(MockRecommendedBookRepository)
	interactor := RecommendedBookInteractor{
		RecommendedBookRepository: repo,
	}
	_, err := interactor.RecommendedBookRepository.All(4)
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	//if len(recommendedBooks) != 4 {
	//	t.Fatalf("Fail expected: 4, got: %v", len(recommendedBooks))
	//}
}
