package repository

import (
	"math3usmartins.com/todo-api-golang/account/repository"
	"math3usmartins.com/todo-api-golang/database/mysql"
)

func NewRepository(adapter mysql.Adapter) (repository.MysqlRepository, error) {
	repo := repository.MysqlRepository{}

	err := repo.WithAdapter(adapter)

	return repo, err
}
