package account

type User struct {
	Uuid     string
	Email    string
	Password string
	Roles    []string
}
