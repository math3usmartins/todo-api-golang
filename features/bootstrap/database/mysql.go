package database

import (
	"math3usmartins.com/todo-api-golang/database/mysql"
)

func NewMysqlAdapter() mysql.Adapter {
	adapter := mysql.Adapter{}

	return adapter
}
