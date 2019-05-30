package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/go-sql-driver/mysql"
	"math3usmartins.com/todo-api-golang/features/bootstrap"
	"os"
)

func FeatureContext(s *godog.Suite) {
	var err error

	dsn := os.Getenv("TODO_API_DATABASE_DSN")

	driver := mysql.MySQLDriver{}

	_, err = driver.Open(dsn)

	if err != nil {
		panic(err)
	}

	adapter := bootstrap.NewDatabase(driver)

	if err != nil {
		panic(err)
	}

	accountContext := bootstrap.AccountContext{}

	err = accountContext.Init(adapter)

	if err != nil {
		panic(err)
	}

	s.Step(
		`^the account service is called to create a user with the following information:$`,
		accountContext.TheAccountServiceIsCalledToCreateAUserWithTheFollowingInformation,
	)

	s.Step(
		`^a user can be found by email "([^"]*)"$`,
		accountContext.AUserCanBeFoundByEmail,
	)
}
