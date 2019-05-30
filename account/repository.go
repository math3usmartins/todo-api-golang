package account

type Repository interface {
	Create(user User) error
	Update(uuid string, user User) error
	Delete(uuid string) error
	Get(uuid string) (User, error)
	FindByEmail(email string) (User, error)
}
