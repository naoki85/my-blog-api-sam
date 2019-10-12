package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"testing"
)

type MockPostRepository struct {
}

func (repo *MockPostRepository) All() (model.Posts, int, error) {
	posts := model.Posts{
		model.Post{Id: 1, PublishedAt: "2019-01-01 00:00:00"},
		model.Post{Id: 2, PublishedAt: "2019-02-01 00:00:00"},
		model.Post{Id: 3, PublishedAt: "2019-03-01 00:00:00"},
		model.Post{Id: 4, PublishedAt: "2019-03-01 00:00:00"},
		model.Post{Id: 5, PublishedAt: "2019-03-01 00:00:00"},
		model.Post{Id: 6, PublishedAt: "2019-03-01 00:00:00"},
		model.Post{Id: 7, PublishedAt: "2019-03-01 00:00:00"},
		model.Post{Id: 8, PublishedAt: "2019-03-01 00:00:00"},
		model.Post{Id: 9, PublishedAt: "2019-03-01 00:00:00"},
		model.Post{Id: 10, PublishedAt: "2019-03-01 00:00:00"},
		model.Post{Id: 11, PublishedAt: "2019-03-01 00:00:00"},
		model.Post{Id: 11, PublishedAt: "2021-01-01 00:00:00"},
	}
	return posts, 11, nil
}

func (repo *MockPostRepository) FindById(int) (model.Post, error) {
	post := model.Post{
		Id: 1,
	}
	return post, nil
}

type MockIdCounterRepository struct{}

func (repo *MockIdCounterRepository) FindMaxIdByIdentifier(i string) (int, error) {
	return 1, nil
}

func (repo *MockIdCounterRepository) UpdateMaxIdByIdentifier(i string, n int) (int, error) {
	return 1, nil
}

func TestShouldPostsIndex(t *testing.T) {
	interactor := initTestPostInteractor()
	posts, count, err := interactor.Index(1)
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if len(posts) != 10 {
		t.Fatalf("Fail expected: 10, got: %v", len(posts))
	}
	if count != 2 {
		t.Fatalf("Fail expected: 2, got: %d", count)
	}

	posts2, _, _ := interactor.Index(2)
	if len(posts2) != 1 {
		t.Fatalf("Fail expected: 1, got: %d", len(posts2))
	}
}

func TestShouldFindPostById(t *testing.T) {
	interactor := initTestPostInteractor()
	post, err := interactor.FindById(1)
	if err != nil {
		t.Fatalf("Cannot get recommended_books: %s", err)
	}
	if post.Id != 1 {
		t.Fatalf("Fail expected id: 1, got: %v", post)
	}
}

func initTestPostInteractor() PostInteractor {
	return PostInteractor{
		PostRepository:      new(MockPostRepository),
		IdCounterRepository: new(MockIdCounterRepository),
	}
}
