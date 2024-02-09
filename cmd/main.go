package main

import (
	"context"
	"log"

	"backend/internal/app"
)

// start server
func main() {
	log.Println("Application start!")
	// Создаем контекст
	ctx := context.Background()

	// Создаем Aplication
	application, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Запустите сервер, вызвав метод StartServer у объекта Application
	application.Run()
	log.Println("Application terminated!")
}
