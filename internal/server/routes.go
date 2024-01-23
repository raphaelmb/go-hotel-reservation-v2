package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) Routes() http.Handler {
	// userStore := database.NewUserStore(s.db)
	// userHandler := handlers.NewUserHandler(userStore)
	e := echo.New()

	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
