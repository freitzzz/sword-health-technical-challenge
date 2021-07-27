package main

import (
	"fmt"

	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/data"
	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/logging"
)

func main() {

	db, oerr := data.OpenDbConnection()

	if oerr != nil {

		logging.LogError(fmt.Sprintf("Failure to open DB connection\n%s", oerr))

		panic("Migrations require DB connection")
	}

	merr := data.RunMigration(db)

	if merr != nil {
		logging.LogError(merr.Error())
	}

}
