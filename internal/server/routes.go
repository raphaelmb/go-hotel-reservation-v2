package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routes() http.Handler {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	return e
}
