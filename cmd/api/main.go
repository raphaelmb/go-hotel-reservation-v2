package main

import (
	"log"

	"github.com/raphaelmb/go-hotel-reservation-v2/internal/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Cannot start server")
	}
}
