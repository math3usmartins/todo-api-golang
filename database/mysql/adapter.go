package mysql

import (
	"database/sql/driver"
	"github.com/go-sql-driver/mysql"
)

type Adapter struct {
	driver     mysql.MySQLDriver
	connection driver.Conn
}

func (adapter Adapter) Connect(dsn string) error {
	connection, err := adapter.driver.Open(dsn)

	if err != nil {
		return err
	}

	adapter.connection = connection

	defer adapter.connection.Close()

	return nil
}

func (adapter Adapter) WithDriver(driver mysql.MySQLDriver) {
	adapter.driver = driver
}

func (adapter Adapter) GetDriver() mysql.MySQLDriver {
	return adapter.driver
}
