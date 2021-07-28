package main

import (
	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/data"
	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/http"
	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/logging"
	"github.com/labstack/echo/v4"
)

func main() {

	// Echo instance
	e := echo.New()

	defer e.Close()

	db, oerr := data.OpenDbConnection()

	if oerr != nil {

		logging.LogError("Could not open DB connection on server start")
		logging.LogError(oerr.Error())

		panic("DB connection is required to serve HTTP calls")
	}

	http.RegisterMiddlewares(e, db)

	http.RegisterHandlers(e)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
