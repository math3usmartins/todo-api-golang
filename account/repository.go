package account

import "math3usmartins.com/todo-api-golang/database"

type Repository interface {
	WithDriver(driver database.Driver) error
	Create(account Account) error
	Update(uuid string, account Account) error
	Delete(uuid string) error
	Get(uuid string) (Account, error)
}
