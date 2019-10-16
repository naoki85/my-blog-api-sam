package repository

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"log"
	"os"
	"time"
)

type S3BookrecorderImageRepository struct {
	S3UploadHandler *s3manager.Uploader
	S3Handler       *s3.S3
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

func (repo *S3BookrecorderImageRepository) CreateSignedUrl(filePath string) (string, error) {
	buf := bytes.NewReader(make([]byte, 10*1024*1024))
	h := md5.New()
	_, err := io.Copy(h, buf)
	if err != nil {
		fmt.Println("error creating MD5 checksum", err)
		return "", err
	}

	r, _ := repo.S3Handler.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(repo.bucketName()),
		Key:    aws.String(filePath),
	})
	r.HTTPRequest.Header.Set("Content-MD5", base64.StdEncoding.EncodeToString(h.Sum(nil)))
	url, err := r.Presign(15 * time.Minute)
	if err != nil {
		fmt.Println("error presigning request", err)
		return "", err
	}
	_, err = buf.Seek(0, 0)

	return url, err
}
