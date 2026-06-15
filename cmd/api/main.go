package main

import (
	"log"
	"net/http"

	"github.com/dhrodrigues29/backend-api/internal/config"
	"github.com/dhrodrigues29/backend-api/internal/server"
)

func main() {
	srv := server.New()
	cfg := config.Load()

	log.Printf(
		"server running on %s",
		cfg.ServerAddress,
	)

	if err := http.ListenAndServe(cfg.ServerAddress, srv); err != nil {
		log.Fatal(err)
	}
}