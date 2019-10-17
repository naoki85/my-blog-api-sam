package controller

import (
	"github.com/naoki85/my-blog-api-sam/config"
	"github.com/naoki85/my-blog-api-sam/infrastructure"
	"testing"
)

func TestShouldGetSignedUrl(t *testing.T) {
	config.InitDbConf("")
	c := config.GetDbConf()
	s3Handler, _ := infrastructure.NewS3Handler(c)
	controller := NewImageUploadController(s3Handler)
	_, status := controller.GetSignedUrl("hoge.png")
	if status == config.InternalServerErrorStatus {
		t.Fatal("Could not get signed url")
	}
}
