package transformer_test

import (
	"encoding/json"
	"math3usmartins.com/todo-api-golang/account"
	"math3usmartins.com/todo-api-golang/account/transformer"
	"reflect"
	"testing"
)

func TestJsonTransformer(t *testing.T) {
	t.Run("FromUser", func(t *testing.T) {
		user := account.User{
			"abc-123",
			"someone@email.com",
			"abc123",
			[]string{},
		}

		expectedJsonBytes, _ := json.Marshal(transformer.UserForJson{
			user.Uuid,
			user.Email,
			user.Password,
			user.Roles,
		})

		expectedJson := string(expectedJsonBytes)

		jsonTransformer := transformer.JsonTransformer{}

		actual, err := jsonTransformer.FromUser(user)

		if err != nil {
			t.Error(err)
		}

		if string(expectedJson) != actual {
			t.Errorf("Actual does match expected JSON. Expected '%s'. Got '%s'", expectedJson, actual)
		}
	})

	t.Run("ToUser", func(t *testing.T) {
		givenUserForJson := transformer.UserForJson{
			"abc-123",
			"someone@email.com",
			"abc123",
			[]string{},
		}

		expectedUser := account.User{
			givenUserForJson.Uuid,
			givenUserForJson.Email,
			givenUserForJson.Password,
			givenUserForJson.Roles,
		}

		givenJsonBytes, _ := json.Marshal(givenUserForJson)

		jsonTransformer := transformer.JsonTransformer{}

		actual, err := jsonTransformer.ToUser(string(givenJsonBytes))

		if err != nil {
			t.Error(err)
		}

		if actual.Uuid != expectedUser.Uuid {
			t.Errorf("Actual UUID does not match expected value. Expected '%s'. Got '%s'", expectedUser.Uuid, actual.Uuid)
		}

		if actual.Email != expectedUser.Email {
			t.Errorf("Actual Email does not match expected value. Expected '%s'. Got '%s'", expectedUser.Email, actual.Email)
		}

		if actual.Password != expectedUser.Password {
			t.Errorf("Actual Email does not match expected value. Expected '%s'. Got '%s'", expectedUser.Password, actual.Password)
		}

		if !reflect.DeepEqual(actual.Roles, expectedUser.Roles) {
			t.Errorf("Actual Roles does not match expected value. Expected '%s'. Got '%s'", expectedUser.Roles, actual.Roles)
		}
	})
}
