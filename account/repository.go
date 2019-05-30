package account

type Repository interface {
	Create(account Account) error
	Update(uuid string, account Account) error
	Delete(uuid string) error
	Get(uuid string) (Account, error)
}
