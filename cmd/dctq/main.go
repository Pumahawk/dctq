package main

import (
	"log"

	"github.com/Pumahawk/dctq/internal/controllers"
	"github.com/Pumahawk/dctq/internal/services"
)

func main() {
	log.Println("Starting dctq server.")

	log.Println("Main - Create services.")
	gameService := services.NewGameServiceImpl()
	messageService := services.NewMessageServiceImpl(gameService)

	log.Println("Main - Create server controllers.")
	server := controllers.NewControllerServerImpl(gameService, messageService)

	log.Printf("Main - Start message processor.")
	go messageService.StartServerMessageProcessor()

	log.Println("Main - Start server controllers.")
	log.Fatal(server.ListenAndServe())
}
