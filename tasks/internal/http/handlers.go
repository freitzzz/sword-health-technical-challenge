package http

import (
	"crypto/cipher"
	"fmt"
	"net/http"
	"strconv"

	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/amqp"
	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/data"
	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/domain"
	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/logging"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterHandlers(e *echo.Echo) {

	technicianMiddleware := onlyAllowTechnicianMiddleware()

	e.GET(getTasks, GetTasks)
	e.POST(performTask, PerformTask, technicianMiddleware)
	e.GET(getTask, GetTask)
	e.PUT(updateTask, UpdateTask, technicianMiddleware)
	e.DELETE(deleteTask, DeleteTask)

	echo.NotFoundHandler = useNotFoundHandler()
}

func GetTasks(c echo.Context) error {

	pidx, pperr := ParsePaginationIndex(c.QueryParam(paginationIndex))

	if pperr != nil {
		return InvalidParamBadRequest(c, paginationIndexNotInteger)
	}

	db, uc, _, _, rerr := requestEssentials(c)

	if rerr != nil {
		return rerr
	}

	tasks := getTasksFromDb(c, db, uc, pidx)

	return Ok(c, ToTaskPage(tasks))

}

func PerformTask(c echo.Context) error {

	db, uc, cb, mb, rerr := requestEssentials(c)

	if rerr != nil {
		return rerr
	}

	var tp TaskPerform

	c.Bind(&tp)

	task, nerr := domain.New(uc.ID, tp.Summary, cb)

	if nerr != nil {
		return InvalidParamBadRequest(c, summaryExceeds2500Characters)
	}

	itask, ierr := data.InsertTask(db, task)

	if ierr != nil {
		logging.LogError("Failed to insert task on database after creating it")
		logging.LogError(ierr.Error())

		return InternalServerError(c)
	}

	amqp.PublishNotification(mb, amqp.Notification{
		UserID:  uc.ID,
		Message: fmt.Sprintf("The tech %s performed the task %s on date %s", itask.UserID, domain.Summary(*itask, cb), itask.CreatedAt.String()),
	})

	return Created(c, ToTaskView(*itask, cb))

}

func GetTask(c echo.Context) error {

	db, uc, cb, tid, rerr := requestEssentialsWithTaskID(c)

	if rerr != nil {
		return rerr
	}

	task, qerr := getTaskFromDb(c, db, uc, tid)

	if qerr != nil {
		return qerr
	}

	return Ok(c, ToTaskView(*task, cb))

}

func UpdateTask(c echo.Context) error {

	db, uc, cb, tid, rerr := requestEssentialsWithTaskID(c)

	if rerr != nil {
		return rerr
	}

	task, qerr := getTaskFromDb(c, db, uc, tid)

	if qerr != nil {
		return qerr
	}

	var tu TaskUpdate

	c.Bind(&tu)

	userr := domain.UpdateSummary(task, tu.Summary, cb)

	if userr != nil {
		return InvalidParamBadRequest(c, summaryExceeds2500Characters)
	}

	_, uerr := data.UpdateTask(db, *task)

	if uerr != nil {
		logging.LogError("Failed to update task on database after updating summary")
		logging.LogError(uerr.Error())

		return InternalServerError(c)
	}

	return Ok(c, ToTaskView(*task, cb))

}

func DeleteTask(c echo.Context) error {

	db, uc, _, tid, rerr := requestEssentialsWithTaskID(c)

	if rerr != nil {
		return rerr
	}

	task, qerr := getTaskFromDb(c, db, uc, tid)

	if qerr != nil {
		return qerr
	}

	domain.Disable(task)

	_, uerr := data.UpdateTask(db, *task)

	if uerr != nil {
		logging.LogError("Failed to update task on database after disabling it")
		logging.LogError(uerr.Error())

		return InternalServerError(c)
	}

	return NoContent(c)

}

func useNotFoundHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	}
}

func requestEssentials(c echo.Context) (*gorm.DB, UserContext, cipher.Block, amqp.MailBox, error) {

	db, dok := c.Get(dbMiddlewareKey).(*gorm.DB)
	uc, uok := c.Get(ucMiddlewareKey).(UserContext)
	cb, cok := c.Get(cbMiddlewareKey).(cipher.Block)
	mb, mok := c.Get(mbMiddlewareKey).(amqp.MailBox)

	if !dok {
		logging.LogError("DB not available in middleware")

		return db, uc, cb, mb, InternalServerError(c)
	}

	if !uok {
		logging.LogError("User Context not available in middleware")

		return db, uc, cb, mb, InternalServerError(c)
	}

	if !cok {
		logging.LogError("Cipher Block not available in middleware")

		return db, uc, cb, mb, InternalServerError(c)
	}

	if !mok {
		logging.LogError("Mail Box not available in middleware")

		return db, uc, cb, mb, InternalServerError(c)
	}

	return db, uc, cb, mb, nil

}

func requestEssentialsWithTaskID(c echo.Context) (*gorm.DB, UserContext, cipher.Block, int, error) {

	tid, terr := strconv.Atoi(c.Param(taskId))

	db, uc, cb, _, rerr := requestEssentials(c)

	if rerr != nil {
		return db, uc, cb, tid, rerr
	} else if terr != nil {
		logging.LogError("Task ID parse not successful, middlware allowed it in the first place")
		logging.LogError(terr.Error())

		InternalServerError(c)

		return db, uc, cb, tid, terr
	}

	return db, uc, cb, tid, nil

}

func getTaskFromDb(c echo.Context, db *gorm.DB, uc UserContext, tid int) (*domain.Task, error) {
	var task *domain.Task
	var qerr error

	if IsTechnician(uc) {
		task, qerr = data.QueryUserTaskById(db, uc.ID, tid)
	} else {
		task, qerr = data.QueryTaskById(db, tid)
	}

	if qerr != nil {
		logging.LogWarning(fmt.Sprintf("User %s with role %d tried to access task %d, but task was not found", uc.ID, uc.Role, tid))
		logging.LogError(qerr.Error())

		NotFound(c)

		return task, qerr
	} else {
		return task, nil
	}
}

func getTasksFromDb(c echo.Context, db *gorm.DB, uc UserContext, pidx int) []*domain.Task {
	var tasks []*domain.Task

	if IsTechnician(uc) {
		tasks = data.QueryUserTasks(db, uc.ID, pidx)
	} else {
		tasks = data.QueryTasks(db, pidx)
	}

	return tasks

}
