package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/naoki85/my-blog-api-sam/config"
)

func NewS3UploaderHandler(c *config.Config) (*s3manager.Uploader, error) {
	var awsConf *aws.Config
	if len(c.DynamoDbEndpoint) > 0 {
		awsConf = &aws.Config{
			Credentials:      credentials.NewStaticCredentials("hogehoge", "fugafuga", ""),
			Region:           aws.String("ap-northeast-1"),
			Endpoint:         aws.String(c.S3Endpoint),
			S3ForcePathStyle: aws.Bool(true),
		}
	} else {
		awsConf = &aws.Config{
			Region: aws.String("ap-northeast-1"),
		}
	}
	sess := session.Must(session.NewSession(awsConf))
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		// Define a strategy that will buffer 25 MiB in memory
		u.BufferProvider = s3manager.NewBufferedReadSeekerWriteToPool(25 * 1024 * 1024)
	})
	return uploader, nil
}
