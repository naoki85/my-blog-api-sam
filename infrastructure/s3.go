package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/naoki85/my-blog-api-sam/config"
	_interface "github.com/naoki85/my-blog-api-sam/interface"
	"time"
)

type S3Handler struct {
	S3 *s3.S3
}

func (h S3Handler) CreateSignedUrl(input _interface.S3Input) (string, error) {
	r, _ := h.S3.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(input.Bucket),
		Key:    aws.String(input.Key),
	})

	url, err := r.Presign(15 * time.Minute)
	return url, err
}

func NewS3Handler(c *config.Config) (S3Handler, error) {
	var awsConf *aws.Config
	if len(c.S3Endpoint) > 0 {
		awsConf = &aws.Config{
			Credentials:      credentials.NewStaticCredentials("hogehoge", "fugafuga", ""),
			Region:           aws.String("ap-northeast-1"),
			Endpoint:         aws.String(c.S3Endpoint),
			DisableSSL:       aws.Bool(true),
			S3ForcePathStyle: aws.Bool(true),
		}
	} else {
		awsConf = &aws.Config{
			Region: aws.String("ap-northeast-1"),
		}
	}
	sess := session.Must(session.NewSession(awsConf))
	h := S3Handler{S3: s3.New(sess)}
	return h, nil
}
