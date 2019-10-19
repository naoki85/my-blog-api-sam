package usecase

import (
	"log"
)

type ImageUploadInteractor struct {
	S3BookrecorderRepository S3BookrecorderImageRepository
}

func (interactor *ImageUploadInteractor) GetSignedUrl(fileName string) (url string, err error) {
	url, err = interactor.S3BookrecorderRepository.CreateSignedUrl(fileName)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return
}
