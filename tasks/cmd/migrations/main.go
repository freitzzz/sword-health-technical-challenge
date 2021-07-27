package main

import (
	"fmt"

	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/data"
)

func main() {

	db, oerr := data.OpenDbConnection()

	if oerr != nil {
		panic(fmt.Sprintf("Failure to open DB connection\n%s", oerr))
	}

	data.RunMigration(db)

}
