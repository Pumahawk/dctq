package controllers

import (
	"net/http"

	"github.com/Pumahawk/dctq/internal/services"
)

const (
	StatusEndpoint     = "/status"
	StatusByIdEndpoint = "/status/{id}"
	MessagesEndpoint   = "/status/{id}/messages"
)

type ControllerServer interface {
	ListenAndServe()
}

type ControllerServerImpl struct {
	ServerMux *http.ServeMux
}

func NewControllerServerImpl(statusService services.StatusService, messageService services.MessageService) *ControllerServerImpl {
	statusController := NewStatusController(statusService)
	messagesController := NewMessagesController(messageService, statusService)

	serverMux := http.NewServeMux()
	serverMux.Handle("GET "+StatusEndpoint, statusController.GetAll())
	serverMux.Handle("POST "+StatusEndpoint, statusController.Create())
	serverMux.Handle("GET "+StatusByIdEndpoint, statusController.GetById())
	serverMux.Handle("PUT "+StatusByIdEndpoint, statusController.Update())
	serverMux.Handle("GET "+MessagesEndpoint, messagesController.Follow())
	serverMux.Handle("POST "+MessagesEndpoint, messagesController.Send())

	return &ControllerServerImpl{
		ServerMux: serverMux,
	}
}

func (c *ControllerServerImpl) ListenAndServe() error {
	return http.ListenAndServe(":8080", c.ServerMux)
}
