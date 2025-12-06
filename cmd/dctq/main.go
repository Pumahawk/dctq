package main

import (
	"log"
	"net/http"

	"github.com/Pumahawk/dctq/internal/controllers"
	"github.com/Pumahawk/dctq/internal/services"
)

func main() {
	log.Println("Starting dctq server.")
	gameService := services.NewGameServiceImpl()
	messageService := services.NewMessageServiceImpl(gameService)
	gameController := controllers.NewGamesController(gameService)
	messagesController := *controllers.NewMessagesController(messageService)
	http.Handle("GET "+controllers.GamesEndpoint, gameController.GetAll())
	http.Handle("POST "+controllers.GamesEndpoint, gameController.Create())
	http.Handle("GET "+controllers.GameByIdEndpoint, gameController.GetById())
	http.Handle("PUT "+controllers.GameByIdEndpoint, gameController.Update())
	http.Handle("GET "+controllers.MessagesEndpoint, messagesController.Follow())
	http.Handle("POST "+controllers.MessagesEndpoint, messagesController.Send())
	go messageService.StartServerMessageProcessor()
	log.Println("Start Cluedo server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
