package http

import (
	"fmt"
	"net/http"
	"strconv"

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

	db, uc, rerr := requestEssentials(c)

	if rerr != nil {
		return rerr
	}

	tasks := getTasksFromDb(c, db, uc, pidx)

	return Ok(c, ToTaskPage(tasks))

}

func PerformTask(c echo.Context) error {

	db, uc, rerr := requestEssentials(c)

	if rerr != nil {
		return rerr
	}

	var tp TaskPerform

	c.Bind(&tp)

	task, nerr := domain.New(uc.ID, tp.Summary)

	if nerr != nil {
		return InvalidParamBadRequest(c, summaryExceeds2500Characters)
	}

	_, ierr := data.InsertTask(db, task)

	if ierr != nil {
		logging.LogError("Failed to insert task on database after creating it")
		logging.LogError(ierr.Error())

		return InternalServerError(c)
	}

	return Created(c, ToTaskView(task))

}

func GetTask(c echo.Context) error {

	db, uc, tid, rerr := requestEssentialsWithTaskID(c)

	if rerr != nil {
		return rerr
	}

	task, qerr := getTaskFromDb(c, db, uc, tid)

	if qerr != nil {
		return qerr
	}

	return Ok(c, ToTaskView(*task))

}

func UpdateTask(c echo.Context) error {

	db, uc, tid, rerr := requestEssentialsWithTaskID(c)

	if rerr != nil {
		return rerr
	}

	task, qerr := getTaskFromDb(c, db, uc, tid)

	if qerr != nil {
		return qerr
	}

	var tu TaskUpdate

	c.Bind(&tu)

	userr := domain.UpdateSummary(task, tu.Summary)

	if userr != nil {
		return InvalidParamBadRequest(c, summaryExceeds2500Characters)
	}

	_, uerr := data.UpdateTask(db, *task)

	if uerr != nil {
		logging.LogError("Failed to update task on database after updating summary")
		logging.LogError(uerr.Error())

		return InternalServerError(c)
	}

	return Ok(c, ToTaskView(*task))

}

func DeleteTask(c echo.Context) error {

	db, uc, tid, rerr := requestEssentialsWithTaskID(c)

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

func requestEssentials(c echo.Context) (*gorm.DB, UserContext, error) {

	db, dok := c.Get(dbMiddlewareKey).(*gorm.DB)
	uc, uok := c.Get(ucMiddlewareKey).(UserContext)

	if !dok {
		logging.LogError("DB not available in DeleteTask middleware")

		return db, uc, InternalServerError(c)
	}

	if !uok {
		logging.LogError("User Context not available in DeleteTask middleware")

		return db, uc, InternalServerError(c)
	}

	return db, uc, nil

}

func requestEssentialsWithTaskID(c echo.Context) (*gorm.DB, UserContext, int, error) {

	tid, terr := strconv.Atoi(c.Param(taskId))

	db, uc, rerr := requestEssentials(c)

	if rerr != nil {
		return db, uc, tid, rerr
	} else if terr != nil {
		logging.LogError("Task ID parse not successful, middlware allowed it in the first place")
		logging.LogError(terr.Error())

		return db, uc, tid, InternalServerError(c)
	}

	return db, uc, tid, nil

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
		return task, NotFound(c)
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
