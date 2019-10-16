package infrastructure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/naoki85/my-blog-api-sam/config"
)

func NewS3Handler(c *config.Config) (*s3.S3, error) {
	var awsConf *aws.Config
	if len(c.S3Endpoint) > 0 {
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
	return s3.New(sess), nil
}
