package mysql

type MysqlAdapter struct {
	connection interface{}
}

func (adapter MysqlAdapter) Connect(dsn string) error {
	return nil
}
