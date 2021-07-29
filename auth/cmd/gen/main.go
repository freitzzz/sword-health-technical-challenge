package main

import (
	"fmt"

	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/data"
	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/domain"
	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/logging"
)

func main() {

	db, oerr := data.OpenDbConnection()

	if oerr != nil {

		logging.LogError(fmt.Sprintf("Failure to open DB connection\n%s", oerr))

		panic("Generate fake data requires DB connection")
	}

	_, merr := data.InsertUser(db, domain.MockTechnician)
	_, merr = data.InsertUser(db, domain.MockManager)

	if merr != nil {
		logging.LogError(merr.Error())
	}

}
