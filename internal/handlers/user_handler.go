package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raphaelmb/go-hotel-reservation-v2/internal/database"
)

type UserHandler struct {
	userStore database.UserStore
}

func NewUserHandler(userStore *database.UserStore) *UserHandler {
	return &UserHandler{
		userStore: *userStore,
	}
}

func (h *UserHandler) HandleGetUser(c echo.Context) error {
	r, err := h.userStore.GetUsers(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, r)
}
