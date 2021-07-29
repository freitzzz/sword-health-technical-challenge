package http

import (
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
