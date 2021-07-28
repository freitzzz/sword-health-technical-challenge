package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TestDbAccessMiddlewareProvidesDbInstanceUsingDbMiddlewareKey(t *testing.T) {
	e := echo.New()

	db := gorm.DB{}

	e.Use(dbAccessMiddleware(&db))

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	c := e.AcquireContext()

	pdb, ok := c.Get(dbMiddlewareKey).(*gorm.DB)

	if !ok {
		t.Fatalf("DB Middleware was registered but when accessing the DB, it was not provided")

	}

	if &db != pdb {
		t.Fatalf("Provided DB in middleware is different than the one provided to the middleware")
	}

}

func TestOnlyAllowManagerAccessMiddlewareReturnsInternalServerErrorIfUserContextIsNotPresent(t *testing.T) {
	e := echo.New()

	e.Use(onlyAllowManagerMiddleware())

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("User Context was not present and should respond with internal server error, but replied %d instead", rec.Code)

	}

}

func TestOnlyAllowManagerAccessMiddlewareReturnsUnauthorizedIfIsTechnicianCallReturnsFalse(t *testing.T) {
	e := echo.New()

	uc := UserContext{ID: "x", Role: 1}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			c.Set(ucMiddlewareKey, uc)
			next(c)

			return nil
		}
	})

	e.Use(onlyAllowManagerMiddleware())

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("User Context is present, but user is not technician, so it should respond with unauthorized, but replied %d instead", rec.Code)

	}

}

func TestOnlyAllowManagerAccessMiddlewareProceedsMiddlewareIfIsTechnicianCallReturnsTrue(t *testing.T) {
	e := echo.New()

	uc := UserContext{ID: "x", Role: 0}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			c.Set(ucMiddlewareKey, uc)
			next(c)

			return nil
		}
	})

	e.Use(onlyAllowManagerMiddleware())

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("User Context is present, user is technician, so it should respond with ok, but replied %d instead", rec.Code)

	}

}

func TestResourceIdentifierValidationMiddlewareProceedsMiddlewareIfNoNotificationResourceIdentifierExistsInPathParameters(t *testing.T) {
	e := echo.New()

	e.Use(resourceIdentifierValidationMiddleware())

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("No resource identifier exists in path parameters, so it should respond with ok, but replied %d instead", rec.Code)
	}

}

func TestResourceIdentifierValidationMiddlewareReturnsNotFoundIfNotificationResourceIdentifierExistsInPathParametersButIsEmpty(t *testing.T) {
	e := echo.New()

	e.GET(fmt.Sprintf("/:%s", notificationId), func(c echo.Context) error {
		return c.NoContent(0)
	})

	e.Use(resourceIdentifierValidationMiddleware())

	req := httptest.NewRequest(echo.GET, "/:id", nil)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("An empty resource identifier exists in path parameters, so it should respond with not found, but replied %d instead", rec.Code)
	}

}

func TestResourceIdentifierValidationMiddlewareReturnsNotFoundIfNotificationResourceIdentifierExistsInPathParametersButItsValueIsNegative(t *testing.T) {
	e := echo.New()

	e.GET(fmt.Sprintf("/:%s", notificationId), func(c echo.Context) error {
		return c.NoContent(0)
	})

	e.Use(resourceIdentifierValidationMiddleware())

	req := httptest.NewRequest(echo.GET, "/-2", nil)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("A negative resource identifier exists in path parameters, so it should respond with not found, but replied %d instead", rec.Code)
	}

}

func TestTranslateHeadersInUserContextMiddlewareReturnsUnauthorizedIfUserIdIsNotPresentInHeaders(t *testing.T) {
	e := echo.New()

	e.Use(translateHeadersInUserContextMiddleware())

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("No user id was passed on headers, so it should respond with unauthorized, but replied %d instead", rec.Code)
	}

}

func TestTranslateHeadersInUserContextMiddlewareReturnsUnauthorizedIfUserRoleIsNotPresentInHeaders(t *testing.T) {
	e := echo.New()

	e.Use(translateHeadersInUserContextMiddleware())

	req := httptest.NewRequest(echo.GET, "/", nil)

	req.Header.Set(userIdHeader, "x")

	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("No user role was passed on headers, so it should respond with unauthorized, but replied %d instead", rec.Code)
	}

}

func TestTranslateHeadersInUserContextMiddlewareReturnsUnauthorizedIfUserRoleValueIsLowerThan0(t *testing.T) {
	e := echo.New()

	e.Use(translateHeadersInUserContextMiddleware())

	req := httptest.NewRequest(echo.GET, "/", nil)

	req.Header.Set(userIdHeader, "x")
	req.Header.Set(roleHeader, "-1")

	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("An user role was passed on headers, and its value is -1 so it should respond with unauthorized, but replied %d instead", rec.Code)
	}

}

func TestTranslateHeadersInUserContextMiddlewareReturnsUnauthorizedIfUserRoleValueIsHigherThan1(t *testing.T) {
	e := echo.New()

	e.Use(translateHeadersInUserContextMiddleware())

	req := httptest.NewRequest(echo.GET, "/", nil)

	req.Header.Set(userIdHeader, "x")
	req.Header.Set(roleHeader, "2")

	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("An user role was passed on headers, and its value is 2 so it should respond with unauthorized, but replied %d instead", rec.Code)
	}

}

func TestTranslateHeadersInUserContextMiddlewareSetsUserContextIfUserRoleValueIs0(t *testing.T) {
	e := echo.New()

	e.Use(translateHeadersInUserContextMiddleware())

	req := httptest.NewRequest(echo.GET, "/", nil)

	req.Header.Set(userIdHeader, "x")
	req.Header.Set(roleHeader, "0")

	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	c := e.AcquireContext()

	_, ucok := c.Get(ucMiddlewareKey).(UserContext)

	if !ucok {
		t.Fatalf("An user role with value 0 was passed on headers, so it should set user context on the request, but it was not provided")
	}

}

func TestTranslateHeadersInUserContextMiddlewareSetsUserContextIfUserRoleValueIs1(t *testing.T) {
	e := echo.New()

	e.Use(translateHeadersInUserContextMiddleware())

	req := httptest.NewRequest(echo.GET, "/", nil)

	req.Header.Set(userIdHeader, "x")
	req.Header.Set(roleHeader, "1")

	rec := httptest.NewRecorder()
	e.NewContext(req, rec)
	e.ServeHTTP(rec, req)

	c := e.AcquireContext()

	_, ucok := c.Get(ucMiddlewareKey).(UserContext)

	if !ucok {
		t.Fatalf("An user role with value 1 was passed on headers, so it should set user context on the request, but it was not provided")
	}

}
