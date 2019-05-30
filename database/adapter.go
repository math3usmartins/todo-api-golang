package database

type Adapter interface {
	Connect(dsn string) error
}
