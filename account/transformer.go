package account

type UserTransformer interface {
	FromUser(user User) (interface{}, error)
	ToUser(data interface{}) (User, error)
}
