package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
	"io"
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
		model.Post{Id: 12, PublishedAt: "2021-01-01 00:00:00"},
	}
	return posts, 11, nil
}

func (repo *MockPostRepository) FindById(int) (model.Post, error) {
	post := model.Post{
		Id: 1,
	}
	return post, nil
}

func (repo *MockPostRepository) Create(params repository.PostCreateParams) error {
	return nil
}

type MockIdCounterRepository struct{}

func (repo *MockIdCounterRepository) FindMaxIdByIdentifier(i string) (int, error) {
	return 1, nil
}

func (repo *MockIdCounterRepository) UpdateMaxIdByIdentifier(i string, n int) (int, error) {
	return 1, nil
}

type MockS3BookrecorderImageRepository struct{}

func (repo *MockS3BookrecorderImageRepository) Create(f string, b io.Reader) error {
	return nil
}

func TestShouldPostsIndex(t *testing.T) {
	interactor := initTestPostInteractor()

	t.Run("Successful Request", func(t *testing.T) {
		posts, count, err := interactor.Index(1, false)
		if err != nil {
			t.Fatalf("Cannot get post: %s", err)
		}
		if len(posts) != 10 {
			t.Fatalf("Fail expected: 10, got: %v", len(posts))
		}
		if count != 2 {
			t.Fatalf("Fail expected: 2, got: %d", count)
		}

		posts2, _, _ := interactor.Index(2, false)
		if len(posts2) != 1 {
			t.Fatalf("Fail expected: 1, got: %d", len(posts2))
		}
	})

	t.Run("Admin request", func(t *testing.T) {
		posts, _, _ := interactor.Index(2, true)
		if len(posts) != 2 {
			t.Fatalf("Fail expected: 2, got: %d", len(posts))
		}
	})
}

func TestShouldFindPostById(t *testing.T) {
	interactor := initTestPostInteractor()
	post, err := interactor.FindById(1)
	if err != nil {
		t.Fatalf("Cannot get post: %s", err)
	}
	if post.Id != 1 {
		t.Fatalf("Fail expected id: 1, got: %v", post)
	}
}

func TestShouldCreatePost(t *testing.T) {
	interactor := initTestPostInteractor()
	params := PostInteractorCreateParams{
		Category:    "aws",
		Title:       "Test title",
		Content:     "Test content",
		ImageUrl:    "test.png",
		Active:      "published",
		PublishedAt: "2019-10-01 00:00:00",
	}

	err := interactor.Create(params)
	if err != nil {
		t.Fatalf("Could not create recommended book: %s", err.Error())
	}
}

func initTestPostInteractor() PostInteractor {
	return PostInteractor{
		PostRepository:                new(MockPostRepository),
		IdCounterRepository:           new(MockIdCounterRepository),
		S3BookrecorderImageRepository: new(MockS3BookrecorderImageRepository),
	}
}
