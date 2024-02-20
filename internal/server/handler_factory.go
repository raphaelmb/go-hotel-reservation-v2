package server

import (
	"github.com/raphaelmb/go-hotel-reservation-v2/internal/database"
	"github.com/raphaelmb/go-hotel-reservation-v2/internal/handlers"
)

func UserHandler(db *database.Service) *handlers.UserHandler {
	userStore := database.NewUserStore(db)
	return handlers.NewUserHandler(&userStore)
}
