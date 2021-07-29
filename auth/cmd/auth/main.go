package main

import (
	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/data"
	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/http"
	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/logging"
	"github.com/labstack/echo/v4"
)

func main() {

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

	e.Logger.Fatal(e.Start(http.ServerAddress()))
}
