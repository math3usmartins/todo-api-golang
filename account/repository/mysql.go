package repository

import (
	"math3usmartins.com/todo-api-golang/account"
	"math3usmartins.com/todo-api-golang/database"
)

type MysqlRepository struct {
	driver database.Driver
}

func (repository MysqlRepository) WithDriver(driver database.Driver) error {
	repository.driver = driver

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
