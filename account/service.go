package account

import "errors"

type Service struct {
	repository Repository
}

func (service Service) WithRepository(repository Repository) {
	service.repository = repository
}

func (service Service) SignUp(email string, password string) (User, error) {
	var user User

	return user, errors.New("This method needs to be implemented.")
}
