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

func (accountContext AccountContext) Init(adapter mysql.Adapter) {
	accountRepository := repository.NewRepository(adapter)

	service := account.Service{}

	service.WithRepository(accountRepository)
}

func (accountContext AccountContext) TheAccountServiceIsCalledToCreateAUserWithTheFollowingInformation(dataTable *gherkin.DataTable) error {
	var err error

	row := dataTable.Rows[1]

	_, err = accountContext.service.SignUp(
		row.Cells[0].Value,
		row.Cells[1].Value,
	)

	return err
}

func (accountContext AccountContext) AUserCanBeFoundByEmail(email string) error {
	return godog.ErrPending
}
