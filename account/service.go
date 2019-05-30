package account

type Service struct {
	repository Repository
}

func (service Service) WithRepository(repository Repository) error {
	service.repository = repository

	return nil
}

func (service Service) SignUp(email string, password string) {

}

func (service Service) SignIn(email string, password string) {

}
