package main

import (
	"context"
	"log"

	"github.com/Pumahawk/dctq/internal/controllers"
	"github.com/Pumahawk/dctq/internal/services"
)

func main() {
	log.Println("Starting dctq server.")

	log.Println("Main - Create services.")
	statusService := services.NewStatusServiceImpl()
	messageService := services.NewMessageServiceImpl()
	serverMessageProcessorImpl := services.NewServerMessageProcessorImpl(context.TODO(), statusService)

	log.Println("Main - Create server controllers.")
	server := controllers.NewControllerServerImpl(statusService, messageService)

	log.Printf("Main - Start message processor.")
	go serverMessageProcessorImpl.Start()

	log.Println("Main - Start server controllers.")
	log.Fatal(server.ListenAndServe())
}
