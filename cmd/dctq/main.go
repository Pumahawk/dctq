package main

import (
	"log"

	"github.com/Pumahawk/dctq/internal/controllers"
	"github.com/Pumahawk/dctq/internal/services"
)

func main() {
	log.Println("Starting dctq server.")

	log.Println("Main - Create services.")
	statusService := services.NewStatusServiceImpl()
	messageService := services.NewMessageServiceImpl(statusService)

	log.Println("Main - Create server controllers.")
	server := controllers.NewControllerServerImpl(statusService, messageService)

	log.Printf("Main - Start message processor.")
	go messageService.StartServerMessageProcessor()

	log.Println("Main - Start server controllers.")
	log.Fatal(server.ListenAndServe())
}
