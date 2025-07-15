package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)

	srv := server.ServerFunc(logger)

	logger.Printf("Server starting on http://localhost%s\n", srv.HttpServer.Addr)
	if err := srv.HttpServer.ListenAndServe(); err != nil {
		logger.Fatalf("Server failed: %v", err)
	}
}
