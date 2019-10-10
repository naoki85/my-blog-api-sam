package usecase

import (
	"github.com/naoki85/my-blog-api-sam/model"
	"github.com/naoki85/my-blog-api-sam/repository"
	"log"
)

type RecommendedBookInteractor struct {
	RecommendedBookRepository RecommendedBookRepository
	IdCounterRepository       IdCounterRepository
}

type RecommendedBookInteractorCreateParams struct {
	Link      string `json:"link"`
	ImageUrl  string `json:"imageUrl"`
	ButtonUrl string `json:"buttonUrl"`
}

func (interactor *RecommendedBookInteractor) All(limit int) (recommendedBooks model.RecommendedBooks, err error) {
	results, err := interactor.RecommendedBookRepository.All()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	count, err := interactor.IdCounterRepository.FindMaxIdByIdentifier("RecommendedBooks")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	var minId int
	if count-limit > 0 {
		minId = count - limit
	} else {
		minId = 0
	}
	for _, book := range results {
		if book.Id > minId {
			recommendedBooks = append(recommendedBooks, book)
		}
	}

	return
}

func (interactor *RecommendedBookInteractor) Create(params RecommendedBookInteractorCreateParams) (err error) {
	maxId, err := interactor.IdCounterRepository.FindMaxIdByIdentifier("RecommendedBooks")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	newId := maxId + 1
	_, err = interactor.IdCounterRepository.UpdateMaxIdByIdentifier("RecommendedBooks", newId)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	var inputParams = repository.RecommendedBookCreateParams{
		Id:        newId,
		Link:      params.Link,
		ImageUrl:  params.ImageUrl,
		ButtonUrl: params.ButtonUrl,
	}
	err = interactor.RecommendedBookRepository.Create(inputParams)
	return err
}
