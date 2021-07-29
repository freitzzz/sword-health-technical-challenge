package http

import (
	"net/http"

	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/data"
	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/domain"
	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/logging"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterHandlers(e *echo.Echo) {

	e.POST(authenticate, Authenticate)

	echo.NotFoundHandler = useNotFoundHandler()
}

func Authenticate(c echo.Context) error {

	db, jb, rerr := requestEssentials(c)

	if rerr != nil {
		return rerr
	}

	var ua UserAuth

	c.Bind(&ua)

	u, qerr := data.QueryUserByIdentifier(db, ua.Identifier)

	if qerr != nil {
		return NotAuthorized(c)
	}

	if !domain.ValidAuth(*u, ua.Identifier, ua.Secret) {
		return NotAuthorized(c)
	}

	us, serr := domain.NewSession(*u, jb)

	if serr != nil {
		return NotAuthorized(c)
	}

	_, ierr := data.InsertUserSession(db, us)

	if ierr != nil {
		logging.LogError("Failed to insert user session on database after creating it")
		logging.LogError(ierr.Error())

		return NotAuthorized(c)
	}

	c.Response().Header().Add("Authorization: Bearer", us.Token)

	return Ok(c)

}

func useNotFoundHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	}
}

func requestEssentials(c echo.Context) (*gorm.DB, domain.JWTBundle, error) {

	db, dok := c.Get(dbMiddlewareKey).(*gorm.DB)

	jb, lok := c.Get(jbMiddlewareKey).(domain.JWTBundle)

	if !dok {
		logging.LogError("DB not available in middleware")

		return db, jb, InternalServerError(c)
	}

	if !lok {
		logging.LogError("JWT Bundle not available in middleware")

		return db, jb, InternalServerError(c)
	}

	return db, jb, nil

}
