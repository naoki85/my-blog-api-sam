package repository

import _interface "github.com/naoki85/my-blog-api-sam/interface"

type S3Handler interface {
	CreateSignedUrl(_interface.S3Input) (string, error)
}
