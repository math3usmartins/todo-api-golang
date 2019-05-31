package mysql

import (
	"errors"
	"math3usmartins.com/todo-api-golang/account"
	"math3usmartins.com/todo-api-golang/database/mysql"
)

type Repository struct {
	adapter mysql.Adapter
}

func (repository Repository) WithAdapter(adapter mysql.Adapter) {
	repository.adapter = adapter
}

func (repository Repository) Create(user account.User) error {
	return errors.New("This method needs to be implemented.")
}

func (repository Repository) Update(uuid string, user account.User) error {
	return errors.New("This method needs to be implemented.")
}

func (repository Repository) Delete(uuid string) error {
	return errors.New("This method needs to be implemented.")
}

func (repository Repository) Get(uuid string) (account.User, error) {
	var user account.User
	var err error

	err = errors.New("This method needs to be implemented.")

	return user, err
}

func (repository Repository) FindByEmail(email string) (account.User, error) {
	var user account.User
	var err error

	err = errors.New("This method needs to be implemented.")

	return user, err
}
