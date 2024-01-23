package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/raphaelmb/go-hotel-reservation-v2/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   *database.Service
}

func NewServer() (*http.Server, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("Error: could not parse PORT env variable")
	}
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}
	return &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}, nil
}
