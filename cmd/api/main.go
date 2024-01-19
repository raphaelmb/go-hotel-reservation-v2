package main

import (
	"log"

	"github.com/raphaelmb/go-hotel-reservation-v2/internal/server"
)

func main() {
	server := server.NewServer()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Cannot start server")
	}
}
