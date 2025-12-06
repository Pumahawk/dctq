package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pumahawk/dctq/internal/dto"
	"github.com/Pumahawk/dctq/internal/mappers"
	"github.com/Pumahawk/dctq/internal/services"
)

type MessagesController struct {
	messageService services.MessageService
}

func NewMessagesController(messageService services.MessageService) *MessagesController {
	return &MessagesController{messageService: messageService}
}

func (c *MessagesController) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		dto := dto.SendMessageRequestDto{}
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			log.Printf("MessagesController Send - Unable to parse body. %s", err)
			w.WriteHeader(500)
			return
		}
		m := mappers.ToCreateMessageModelFromDto(id, &dto)
		err = c.messageService.Send(id, m)
		if err != nil {
			log.Printf("MessagesController Send - Unable to send message. %s", err)
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(201)
	}
}

func (c *MessagesController) Follow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		channel, err := c.messageService.Follow(r.Context(), id)
		if err != nil {
			log.Printf("MessagesController Follow - %s", err)
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.WriteHeader(http.StatusOK)

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming not supported", http.StatusInternalServerError)
			return
		}

		for {
			select {
			case <-r.Context().Done():
				log.Println("Client disconnected")
				return
			case message := <-channel:
				response := dto.FollowMessageResponseDto{
					Type:    message.Type,
					Message: message.Message,
				}
				err = json.NewEncoder(w).Encode(response)
				if err != nil {
					log.Printf("MessagesController Follow - Unable to write response. %s", err)
				}
				flusher.Flush()
			}
		}
	}
}
