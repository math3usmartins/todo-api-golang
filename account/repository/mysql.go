package repository

import (
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
	return nil
}

func (repository MysqlRepository) Update(uuid string, user account.User) error {
	return nil
}

func (repository MysqlRepository) Delete(uuid string) error {
	return nil
}

func (repository MysqlRepository) Get(uuid string) (account.User, error) {
	var user account.User

	return user, nil
}

func (repository MysqlRepository) FindByEmail(email string) (account.User, error) {
	var user account.User

	return user, nil
}
