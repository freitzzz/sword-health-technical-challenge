package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo) {
	e.GET(getTasks, GetTasks)
	e.POST(performTask, PerformTask)
	e.GET(getTask, GetTask)
	e.GET(updateTask, UpdateTask)
	e.GET(deleteTask, DeleteTask)
}

func GetTasks(c echo.Context) error {

	_, err := ParsePaginationIndex(c.QueryParam(paginationIndex))

	if err != nil {
		return InvalidParamBadRequest(c, paginationIndexNotInteger)
	}

	return c.String(http.StatusOK, "")

}

func PerformTask(c echo.Context) error {

	var tp TaskPerform

	c.Bind(&tp)

	return c.String(http.StatusOK, "")

}

func GetTask(c echo.Context) error {

	return c.String(http.StatusOK, "")

}

func UpdateTask(c echo.Context) error {

	tid := c.Param(taskId)

	return c.String(http.StatusOK, tid)

}

func DeleteTask(c echo.Context) error {

	tid := c.Param(taskId)

	return c.String(http.StatusOK, tid)

}

func InvalidParamBadRequest(c echo.Context, ip InvalidParam) error {
	return c.JSON(http.StatusBadRequest, ip)
}
