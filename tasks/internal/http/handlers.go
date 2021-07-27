package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTasks(c echo.Context) error {

	_, err := ParsePaginationIndex(c.QueryParam(paginationIndex))

	if err != nil {
		return InvalidParamBadRequest(c, paginationIndexNotInteger)
	}

	return c.String(http.StatusOK, "")

}

func CreateTask(c echo.Context) error {

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
