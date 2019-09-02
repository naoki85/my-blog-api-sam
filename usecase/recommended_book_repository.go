package usecase

import "github.com/naoki85/my-blog-api-sam/model"

type RecommendedBookRepository interface {
	All(int) (model.RecommendedBooks, error)
}
