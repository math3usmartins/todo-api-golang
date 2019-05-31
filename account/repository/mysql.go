package repository

import (
	"errors"
	"math3usmartins.com/todo-api-golang/account"
	"math3usmartins.com/todo-api-golang/database/mysql"
)

type MysqlRepository struct {
	adapter mysql.Adapter
}

func (repository MysqlRepository) WithAdapter(adapter mysql.Adapter) {
	repository.adapter = adapter
}

func (repository MysqlRepository) Create(user account.User) error {
	return errors.New("This method needs to be implemented.")
}

func (repository MysqlRepository) Update(uuid string, user account.User) error {
	return errors.New("This method needs to be implemented.")
}

func (repository MysqlRepository) Delete(uuid string) error {
	return errors.New("This method needs to be implemented.")
}

func (repository MysqlRepository) Get(uuid string) (account.User, error) {
	var user account.User
	var err error

	err = errors.New("This method needs to be implemented.")

	return user, err
}

func (repository MysqlRepository) FindByEmail(email string) (account.User, error) {
	var user account.User
	var err error

	err = errors.New("This method needs to be implemented.")

	return user, err
}
