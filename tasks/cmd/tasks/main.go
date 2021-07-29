package main

import (
	aqp "github.com/freitzzz/sword-health-technical-challenge/tasks/internal/amqp"
	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/data"
	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/domain"
	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/http"
	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/logging"
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

	cb, lerr := domain.LoadTaskSummaryCipher()

	if lerr != nil {

		logging.LogError("Failed to load AES Cipher for task summary encryption")
		logging.LogError(lerr.Error())

		panic("Cannot proceed without cipher for encrypting task summary")

	}

	mqc, merr := aqp.OpenMQConnection()

	defer mqc.Close()

	if merr != nil {

		logging.LogError("Could not open AMQP connection on server start")
		logging.LogError(merr.Error())

		panic("AMQP connection is required to consume notification messages")
	}

	mb, serr := aqp.SetupMailBox(mqc)

	if serr != nil {

		logging.LogError("Could not setup mail box")
		logging.LogError(serr.Error())

		panic("Mail Box is required to provide channel and queue in middleware")
	}

	http.RegisterMiddlewares(e, db, cb, mb)

	http.RegisterHandlers(e)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
