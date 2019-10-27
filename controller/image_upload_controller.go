package controller

import (
	"encoding/json"
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/infrastructure"
	"github.com/naoki85/my-blog-api-sam/repository"
	"github.com/naoki85/my-blog-api-sam/usecase"
	"log"
)

type ImageUploadController struct {
	Interactor usecase.ImageUploadInteractor
}

func NewImageUploadController(s3Handler infrastructure.S3Handler) *ImageUploadController {
	return &ImageUploadController{
		Interactor: usecase.ImageUploadInteractor{
			S3BookrecorderRepository: &repository.S3BookrecorderImageRepository{
				S3Handler: s3Handler,
			},
		},
	}
}

func (controller *ImageUploadController) GetSignedUrl(fileName string) ([]byte, int) {
	url, err := controller.Interactor.GetSignedUrl(fileName)
	if err != nil {
		log.Printf("%s", err.Error())
		return []byte{}, config.InternalServerErrorStatus
	}

	data := struct {
		Url string
	}{url}
	resp, err := json.Marshal(data)
	if err != nil {
		log.Printf("%s", err.Error())
		return resp, config.InternalServerErrorStatus
	}
	return resp, config.SuccessStatus
}
