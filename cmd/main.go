package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	fmt.Println("Сервер запущен")
	// Создаем логгер
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)

	// Создаем сервер с помощью ServerFunc
	srv := server.ServerFunc(logger)

	// Запускаем сервер и обрабатываем ошибки
	logger.Printf("Server starting on http://localhost%s\n", srv.HttpServer.Addr)
	if err := srv.HttpServer.ListenAndServe(); err != nil {
		logger.Fatalf("Server failed: %v", err)
	}
}
