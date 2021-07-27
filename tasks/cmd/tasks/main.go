package main

import (
	http "github.com/freitzzz/sword-health-technical-challenge/tasks/internal/http"
	"github.com/labstack/echo/v4"
)

func main() {

	// Echo instance
	e := echo.New()

	http.RegisterHandlers(e)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
