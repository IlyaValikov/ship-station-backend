package main

import (
	"log"

	"backend/internal/api"
)

// start server
func main() {
	log.Println("Application start!")
	api.StartServer()
	log.Println("Application terminated!")
}
