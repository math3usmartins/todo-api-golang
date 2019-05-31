package repository

import (
	"math3usmartins.com/todo-api-golang/account/repository"
	"math3usmartins.com/todo-api-golang/database/mysql"
)

func NewRepository(adapter mysql.Adapter) repository.MysqlRepository {
	repo := repository.MysqlRepository{}
	repo.WithAdapter(adapter)

	return repo
}
