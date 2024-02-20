package database

import (
	"context"

	"github.com/raphaelmb/go-hotel-reservation-v2/internal/models"
)

type UserStore interface {
	GetUsers(ctx context.Context) ([]*models.User, error)
}

type UserStorePg struct {
	db *Service
}

func NewUserStore(db *Service) UserStore {
	return &UserStorePg{
		db: db,
	}
}

func (s *UserStorePg) GetUsers(ctx context.Context) ([]*models.User, error) {
	_, err := s.db.DB.Query("")
	if err != nil {
		return nil, err
	}
	return []*models.User{}, nil
}
