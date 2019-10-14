package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"log"
	"os"
)

type S3BookrecorderImageRepository struct {
	S3UploadHandler *s3manager.Uploader
}

func (repo *S3BookrecorderImageRepository) bucketName() (bucketName string) {
	if bucketName = os.Getenv("S3_BUCKET"); bucketName != "" {
		return bucketName
	} else {
		return "bookrecorder-image"
	}
}

func (repo *S3BookrecorderImageRepository) Create(saveFilePath string, body io.Reader) (err error) {
	_, err = repo.S3UploadHandler.Upload(&s3manager.UploadInput{
		Bucket: aws.String(repo.bucketName()),
		Key:    aws.String(saveFilePath),
		Body:   body,
	})
	if err != nil {
		log.Printf("S3 upload error: %s", err.Error())
		return
	}

	return
}
