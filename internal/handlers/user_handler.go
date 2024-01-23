package handlers

import (
	"github.com/raphaelmb/go-hotel-reservation-v2/internal/database"
)

type UserHandler struct {
	userStore *database.UserStore
}

func NewUserHandler(userStore *database.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

// func (h *UserHandler) HandleGet(c echo.Context) error {
// 	h.userStore.InsertUser()
// 	return nil
// }
