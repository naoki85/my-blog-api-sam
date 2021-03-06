package repository

import (
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/infrastructure"
	"testing"
)

func TestShouldGetSignedUrl(t *testing.T) {
	config.InitDbConf("")
	c := config.GetDbConf()
	s3Handler, _ := infrastructure.NewS3Handler(c)
	repo := S3BookrecorderImageRepository{
		S3Handler: s3Handler,
	}

	_, err := repo.CreateSignedUrl("posts/hoge.png")
	if err != nil {
		t.Fatalf("Cannot create user: %s", err)
	}
}
