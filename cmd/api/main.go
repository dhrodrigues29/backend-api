package main

import (
	"log"
	"net/http"

	"github.com/dhrodrigues29/backend-api/internal/server"
)

func main() {
	srv := server.New()

	log.Println("server running on :8080")

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}