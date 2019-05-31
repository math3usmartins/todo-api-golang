package bootstrap

import (
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"math3usmartins.com/todo-api-golang/account"
	"math3usmartins.com/todo-api-golang/database/mysql"
	"math3usmartins.com/todo-api-golang/features/bootstrap/account/repository"
)

type AccountContext struct {
	service account.Service
}

func (accountContext AccountContext) Init(adapter mysql.Adapter) error {
	var err error

	accountRepository := repository.NewRepository(adapter)

	service := account.Service{}

	err = service.WithRepository(accountRepository)

	if err != nil {
		return err
	}

	return nil
}

func (accountContext AccountContext) TheAccountServiceIsCalledToCreateAUserWithTheFollowingInformation(dataTable *gherkin.DataTable) error {
	row := dataTable.Rows[1]

	accountContext.service.SignUp(
		row.Cells[0].Value,
		row.Cells[1].Value,
	)

	return nil
}

func (accountContext AccountContext) AUserCanBeFoundByEmail(email string) error {
	return godog.ErrPending
}
