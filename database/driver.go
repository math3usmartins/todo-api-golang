package database

type Driver interface {
	Connect(dsn string) error
}
