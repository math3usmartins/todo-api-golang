package transformer

import "math3usmartins.com/todo-api-golang/account"

type JsonTransformer struct{}

func (transformer JsonTransformer) FromUser(user account.User) (interface{}, error) {
	var err error

	json := "{}"

	return json, err
}

func (transformer JsonTransformer) ToUser(data interface{}) (account.User, error) {
	var err error

	user := account.User{}

	return user, err
}
