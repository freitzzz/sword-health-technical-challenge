package http

import (
	"net/http"

	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/logging"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterHandlers(e *echo.Echo) {

	e.POST(authenticate, Authenticate)

	echo.NotFoundHandler = useNotFoundHandler()
}

func Authenticate(c echo.Context) error {

	_, rerr := requestEssentials(c)

	if rerr != nil {
		return rerr
	}

	// var tp TaskPerform

	// c.Bind(&tp)

	// task, nerr := domain.New(uc.ID, tp.Summary, cb)

	// if nerr != nil {
	// 	return InvalidParamBadRequest(c, summaryExceeds2500Characters)
	// }

	// itask, ierr := data.InsertTask(db, task)

	// if ierr != nil {
	// 	logging.LogError("Failed to insert task on database after creating it")
	// 	logging.LogError(ierr.Error())

	// 	return InternalServerError(c)
	// }

	// amqp.PublishNotification(mb, amqp.Notification{
	// 	UserID:  uc.ID,
	// 	Message: fmt.Sprintf("The tech %s performed the task %s on data %s", itask.UserID, domain.Summary(*itask, cb), itask.CreatedAt.String()),
	// })

	return Ok(c)

}

func useNotFoundHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	}
}

func requestEssentials(c echo.Context) (*gorm.DB, error) {

	db, dok := c.Get(dbMiddlewareKey).(*gorm.DB)

	if !dok {
		logging.LogError("DB not available in middleware")

		return db, InternalServerError(c)
	}

	return db, nil

}
