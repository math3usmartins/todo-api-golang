package repository

import (
	mysqlRepository "math3usmartins.com/todo-api-golang/account/repository/mysql"
	"math3usmartins.com/todo-api-golang/database/mysql"
)

func NewRepository(adapter mysql.Adapter) mysqlRepository.Repository {
	repo := mysqlRepository.Repository{}
	repo.WithAdapter(adapter)

	return repo
}
