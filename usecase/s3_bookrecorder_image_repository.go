package usecase

type S3BookrecorderImageRepository interface {
	CreateSignedUrl(string) (string, error)
}
