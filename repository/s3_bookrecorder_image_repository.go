package repository

import (
	"fmt"
	_interface "github.com/naoki85/my-blog-api-sam/interface"
	"os"
	"regexp"
	"strings"
)

type S3BookrecorderImageRepository struct {
	S3Handler S3Handler
}

func (repo *S3BookrecorderImageRepository) bucketName() (bucketName string) {
	if bucketName = os.Getenv("S3_BUCKET"); bucketName != "" {
		return bucketName
	} else {
		return "bookrecorder-image"
	}
}

func (repo *S3BookrecorderImageRepository) CreateSignedUrl(filePath string) (string, error) {
	input := _interface.S3Input{
		Bucket: repo.bucketName(),
		Key:    filePath,
	}

	url, err := repo.S3Handler.CreateSignedUrl(input)
	if err != nil {
		fmt.Println("error presigning request", err)
		return "", err
	}

	// minio の URL が生成される場合は localhost に書き換える
	regEx := regexp.MustCompile(`^http://minio`)
	if regEx.MatchString(url) {
		url = strings.Replace(url, "http://minio", "http://localhost", 1)
	}

	return url, err
}
