package server

import "net/http"

func NewServer() *http.Server {
	return &http.Server{
		Addr:    ":3000",
		Handler: Routes(),
	}
}
