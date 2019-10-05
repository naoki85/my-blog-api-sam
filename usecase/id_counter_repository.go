package usecase

type IdCounterRepository interface {
	FindMaxIdByIdentifier(string) (int, error)
	UpdateMaxIdByIdentifier(string, int) (int, error)
}
