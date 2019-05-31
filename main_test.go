package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/go-sql-driver/mysql"
	"math3usmartins.com/todo-api-golang/features/bootstrap"
	"os"
)

func FeatureContext(s *godog.Suite) {
	var err error

	driver := mysql.MySQLDriver{}

	adapter := bootstrap.NewDatabase(driver)

	err = adapter.Connect(os.Getenv("TODO_API_DATABASE_DSN"))

	if err != nil {
		panic(err)
	}

	accountContext := bootstrap.AccountContext{}

	accountContext.Init(adapter)

	s.Step(
		`^the account service is called to create a user with the following information:$`,
		accountContext.TheAccountServiceIsCalledToCreateAUserWithTheFollowingInformation,
	)

	s.Step(
		`^a user can be found by email "([^"]*)"$`,
		accountContext.AUserCanBeFoundByEmail,
	)
}
