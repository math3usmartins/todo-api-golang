package transformer

import (
	"encoding/json"
	"math3usmartins.com/todo-api-golang/account"
)

type JsonTransformer struct{}

type UserForJson struct {
	Uuid     string   `json:"uuid"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

func (transformer JsonTransformer) FromUser(user account.User) (string, error) {
	var err error

	userForJson := UserForJson{
		user.Uuid,
		user.Email,
		user.Password,
		user.Roles,
	}

	userJson, err := json.Marshal(userForJson)

	return string(userJson), err
}

func (transformer JsonTransformer) ToUser(userJson string) (account.User, error) {
	var err error

	userForJson := UserForJson{}

	json.Unmarshal([]byte(userJson), &userForJson)

	user := account.User{
		userForJson.Uuid,
		userForJson.Email,
		userForJson.Password,
		userForJson.Roles,
	}

	return user, err
}
