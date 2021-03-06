package main

import (
	aqp "github.com/freitzzz/sword-health-technical-challenge/notifications/internal/amqp"
	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/data"
	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/http"
	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/logging"
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

	mqc, merr := aqp.OpenMQConnection()

	defer mqc.Close()

	if merr != nil {

		logging.LogError("Could not open AMQP connection on server start")
		logging.LogError(merr.Error())

		panic("AMQP connection is required to consume notification messages")
	}

	aqp.RegisterHandlers(mqc, db)

	http.RegisterMiddlewares(e, db)

	http.RegisterHandlers(e)

	e.Logger.Fatal(e.Start(http.ServerAddress()))

}
