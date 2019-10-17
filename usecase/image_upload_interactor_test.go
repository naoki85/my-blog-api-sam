package usecase

import (
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/infrastructure"
	"github.com/naoki85/my-blog-api-sam/repository"
	"testing"
)

func TestShouldCreateSignedUrl(t *testing.T) {
	config.InitDbConf("")
	c := config.GetDbConf()
	s3Handler, _ := infrastructure.NewS3Handler(c)
	interactor := ImageUploadInteractor{
		S3BookrecorderRepository: &repository.S3BookrecorderImageRepository{
			S3Handler: s3Handler,
		},
	}

	_, err := interactor.GetSignedUrl("hoge.png")
	if err != nil {
		t.Fatalf("Could not create signed url: %s", err.Error())
	}
}
