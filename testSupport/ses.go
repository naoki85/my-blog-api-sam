package testSupport

func SetupMockSesHandler() MockSesHandler {
	return MockSesHandler{}
}

type MockSesHandler struct{}

func (sesHandler *MockSesHandler) SendMail(to string, title string, body string) error {
	return nil
}
