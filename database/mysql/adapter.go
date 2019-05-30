package mysql

import "github.com/go-sql-driver/mysql"

type Adapter struct {
	driver mysql.MySQLDriver
}

func (adapter Adapter) Connect(dsn string) error {
	return nil
}

func (adapter Adapter) WithDriver(driver mysql.MySQLDriver) error {
	adapter.driver = driver

	return nil
}

func (adapter Adapter) GetDriver() mysql.MySQLDriver {
	return adapter.driver
}
