package usecase

type IdCounterRepository interface {
	FindCountByIdentifier(string) (int, error)
}
