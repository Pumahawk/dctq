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
	statusController := NewStatusController(statusService, messageService)
	messagesController := NewMessagesController(messageService, statusService)

	serverMux := http.NewServeMux()
	serverMux.Handle("GET "+StatusEndpoint, cors(statusController.GetAll()))
	serverMux.Handle("POST "+StatusEndpoint, cors(statusController.Create()))
	serverMux.Handle("GET "+StatusByIdEndpoint, cors(statusController.GetById()))
	serverMux.Handle("POST "+StatusByIdEndpoint, cors(statusController.Update()))
	serverMux.Handle("GET "+MessagesEndpoint, cors(statusController.Follow()))
	serverMux.Handle("POST "+MessagesEndpoint, cors(messagesController.Send()))

	return &ControllerServerImpl{
		ServerMux: serverMux,
	}
}

func (c *ControllerServerImpl) ListenAndServe() error {
	return http.ListenAndServe(":8080", c.ServerMux)
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
