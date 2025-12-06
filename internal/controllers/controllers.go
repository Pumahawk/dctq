package controllers

import (
	"net/http"

	"github.com/Pumahawk/dctq/internal/services"
)

const (
	GamesEndpoint    = "/games"
	GameByIdEndpoint = "/games/{id}"
	MessagesEndpoint = "/games/{id}/messages"
)

type ControllerServer interface {
	ListenAndServe()
}

type ControllerServerImpl struct {
	ServerMux *http.ServeMux
}

func NewControllerServerImpl(gameService services.GameService, messageService services.MessageService) *ControllerServerImpl {
	gameController := NewGamesController(gameService)
	messagesController := NewMessagesController(messageService)

	serverMux := http.NewServeMux()
	serverMux.Handle("GET "+GamesEndpoint, gameController.GetAll())
	serverMux.Handle("POST "+GamesEndpoint, gameController.Create())
	serverMux.Handle("GET "+GameByIdEndpoint, gameController.GetById())
	serverMux.Handle("PUT "+GameByIdEndpoint, gameController.Update())
	serverMux.Handle("GET "+MessagesEndpoint, messagesController.Follow())
	serverMux.Handle("POST "+MessagesEndpoint, messagesController.Send())

	return &ControllerServerImpl{
		ServerMux: serverMux,
	}
}

func (c *ControllerServerImpl) ListenAndServe() error {
	return http.ListenAndServe(":8080", c.ServerMux)
}
