package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) Routes() http.Handler {
	userHandler := UserHandler(s.db)

	e := echo.New()

	e.GET("/health", s.healthHandler)
	e.GET("/users", userHandler.HandleGetUser)

	return e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
