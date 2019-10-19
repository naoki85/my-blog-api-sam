package repository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"regexp"
	"strings"
	"time"
)

type S3BookrecorderImageRepository struct {
	S3Handler *s3.S3
}

func (repo *S3BookrecorderImageRepository) bucketName() (bucketName string) {
	if bucketName = os.Getenv("S3_BUCKET"); bucketName != "" {
		return bucketName
	} else {
		return "bookrecorder-image"
	}
}

func (repo *S3BookrecorderImageRepository) CreateSignedUrl(filePath string) (string, error) {
	r, _ := repo.S3Handler.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(repo.bucketName()),
		Key:    aws.String(filePath),
	})

	url, err := r.Presign(15 * time.Minute)
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
