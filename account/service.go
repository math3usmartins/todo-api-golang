package account

type Service struct {
	repository Repository
}

func (service Service) WithRepository(repository Repository) {
	service.repository = repository
}

func (service Service) SignUp(email string, password string) {

}

func (service Service) SignIn(email string, password string) {

}
