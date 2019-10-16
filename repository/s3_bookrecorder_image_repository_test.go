package repository

import (
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/infrastructure"
	"os"
	"path/filepath"
	"testing"
)

func TestShouldUploadImage(t *testing.T) {
	config.InitDbConf("")
	c := config.GetDbConf()
	s3Uploader, _ := infrastructure.NewS3UploaderHandler(c)
	repo := S3BookrecorderImageRepository{
		S3UploadHandler: s3Uploader,
	}

	path, _ := filepath.Abs("../testSupport/fixture/test.png")
	filename := "test.png"
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		t.Fatalf("Cannot open file: %s", err)
	}

	err = repo.Create(filename, file)
	if err != nil {
		t.Fatalf("Cannot create user: %s", err)
	}
}

func TestShouldGetSignedUrl(t *testing.T) {
	config.InitDbConf("")
	c := config.GetDbConf()
	s3Uploader, _ := infrastructure.NewS3UploaderHandler(c)
	s3Handler, _ := infrastructure.NewS3Handler(c)
	repo := S3BookrecorderImageRepository{
		S3UploadHandler: s3Uploader,
		S3Handler:       s3Handler,
	}

	_, err := repo.CreateSignedUrl("posts/hoge.png")
	if err != nil {
		t.Fatalf("Cannot create user: %s", err)
	}
}
