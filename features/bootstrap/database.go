package bootstrap

import (
	mysqlDriver "github.com/go-sql-driver/mysql"
	"math3usmartins.com/todo-api-golang/database/mysql"
)

func NewDatabase(driver mysqlDriver.MySQLDriver) mysql.Adapter {
	adapter := mysql.Adapter{}
	adapter.WithDriver(driver)

	return adapter
}
