package repository

import (
	"math3usmartins.com/todo-api-golang/account"
	"math3usmartins.com/todo-api-golang/database/mysql"
)

type MysqlRepository struct {
	adapter mysql.Adapter
}

func (repository MysqlRepository) WithAdapter(adapter mysql.Adapter) error {
	repository.adapter = adapter

	return nil
}

func (repository MysqlRepository) Create(account account.Account) error {
	return nil
}

func (repository MysqlRepository) Update(uuid string, account account.Account) error {
	return nil
}

func (repository MysqlRepository) Delete(uuid string) error {
	return nil
}

func (repository MysqlRepository) Get(uuid string) (account.Account, error) {
	var foundAccount account.Account

	return foundAccount, nil
}

func (repository MysqlRepository) FindByEmail(email string) (account.Account, error) {
	var foundAccount account.Account

	return foundAccount, nil
}
