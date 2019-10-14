package usecase

import "io"

type S3BookrecorderImageRepository interface {
	Create(string, io.Reader) error
}
