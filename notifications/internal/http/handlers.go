package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/data"
	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/domain"
	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/logging"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterHandlers(e *echo.Echo) {

	managerMiddleware := onlyAllowManagerMiddleware()

	e.GET(getNotifications, GetNotifications)
	e.DELETE(deleteNotifications, DeleteNotifications)

	echo.NotFoundHandler = useNotFoundHandler()
}

func GetNotifications(c echo.Context) error {

	db, uc, _, rerr := requestEssentials(c)

	if rerr != nil {
		return rerr
	}

	notifications := getNotificationsFromDb(c, db, uc, pidx)

	return Ok(c, ToNotificationPage(notifications))

}

func DeleteNotification(c echo.Context) error {

	db, uc, _, tid, rerr := requestEssentialsWithNotificationID(c)

	if rerr != nil {
		return rerr
	}

	notification, qerr := getNotificationFromDb(c, db, uc, tid)

	if qerr != nil {
		return qerr
	}

	domain.MarkAsRead(notification)

	_, uerr := data.UpdateNotification(db, *notification)

	if uerr != nil {
		logging.LogError("Failed to update notification on database after marking it as read")
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
		logging.LogError("DB not available in middleware")

		return db, uc, InternalServerError(c)
	}

	if !uok {
		logging.LogError("User Context not available in middleware")

		return db, uc, InternalServerError(c)
	}

	return db, uc, nil

}

func requestEssentialsWithNotificationID(c echo.Context) (*gorm.DB, UserContext, int, error) {

	tid, terr := strconv.Atoi(c.Param(notificationId))

	db, uc, rerr := requestEssentials(c)

	if rerr != nil {
		return db, uc, tid, rerr
	} else if terr != nil {
		logging.LogError("Notification ID parse not successful, middlware allowed it in the first place")
		logging.LogError(terr.Error())

		InternalServerError(c)

		return db, uc, tid, terr
	}

	return db, uc, tid, nil

}

func getNotificationFromDb(c echo.Context, db *gorm.DB, uc UserContext, tid int) (*domain.Notification, error) {
	var notification *domain.Notification
	var qerr error

	if IsManager(uc) {
		notification, qerr = data.QueryUserNotificationById(db, uc.ID, tid)
	} else {
		notification, qerr = data.QueryNotificationById(db, tid)
	}

	if qerr != nil {
		logging.LogWarning(fmt.Sprintf("User %s with role %d tried to access notification %d, but notification was not found", uc.ID, uc.Role, tid))
		logging.LogError(qerr.Error())

		NotFound(c)

		return notification, qerr
	} else {
		return notification, nil
	}
}

func getNotificationsFromDb(c echo.Context, db *gorm.DB, uc UserContext, pidx int) []*domain.Notification {
	var notifications []*domain.Notification

	if IsManager(uc) {
		notifications = data.QueryUserNotifications(db, uc.ID, pidx)
	} else {
		notifications = data.QueryNotifications(db, pidx)
	}

	return notifications

}
