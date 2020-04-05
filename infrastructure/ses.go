package infrastructure

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/naoki85/my-blog-api-sam/config"
)

const (
	FromMailAddress = "naoki.yoneyama.85@gmail.com"
)

type SesHandler struct {
	Ses *ses.SES
}

func NewSesHandler(c *config.Config) (SesHandler, error) {
	var awsConf *aws.Config
	awsConf = &aws.Config{
		Region:   aws.String("ap-southeast-2"),
		Endpoint: aws.String(c.SesEndpoint),
	}
	sess := session.Must(session.NewSession(awsConf))
	h := SesHandler{Ses: ses.New(sess)}
	return h, nil
}

func (sesHandler *SesHandler) SendMail(to string, title string, body string) error {
	if len(sesHandler.Ses.Endpoint) < 1 {
		return nil
	}

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(to),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(title),
			},
		},
		Source: aws.String(FromMailAddress),
	}
	_, err := sesHandler.Ses.SendEmail(input)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
