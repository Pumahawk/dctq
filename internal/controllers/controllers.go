package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pumahawk/dctq/internal/dto"
	"github.com/Pumahawk/dctq/internal/mappers"
	"github.com/Pumahawk/dctq/internal/model"
	"github.com/Pumahawk/dctq/internal/services"
)

const (
	GamesEndpoint    = "/games"
	GameByIdEndpoint = "/games/{id}"
	MessagesEndpoint = "/games/{id}/messages"
)

type GamesController struct {
	gameService services.GameService
}

func NewGamesController(gameService services.GameService) *GamesController {
	return &GamesController{
		gameService,
	}
}

type MessagesController struct {
	messageService services.MessageService
}

func NewMessagesController(messageService services.MessageService) *MessagesController {
	return &MessagesController{messageService: messageService}
}

func (c *GamesController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		games, err := c.gameService.GetAll()
		if err != nil {
			log.Printf("GamesController - error. %s", err)
			return
		}
		gamesDto := mappers.ToGetAllGameResponseDto(games)
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(gamesDto)
		if err != nil {
			log.Printf("GamesController - error. %s", err)
			return
		}
	}
}

func (c *GamesController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createGameInfoRequestDto dto.CreateGameInfoRequestDto
		json.NewDecoder(r.Body).Decode(&createGameInfoRequestDto)
		game, err := c.gameService.Create(*mappers.ToSimplGameCreateInfoModel(&createGameInfoRequestDto))
		if err != nil {
			log.Printf("CreateGameController - error. %s", err)
			return
		}
		responseDto := mappers.ToCreateGameInfoResponseDto(game)
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(responseDto)
		if err != nil {
			log.Printf("CreateGameController - error. %s", err)
			return
		}
	}
}

func (c *GamesController) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		gameModel, err := c.gameService.GetById(id)
		if err != nil {
			log.Printf("GetGameByIdController - error. %s", err)
			return
		}
		responseDto := mappers.ToGetGameByIdResponseDto(gameModel)
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(responseDto)
		if err != nil {
			log.Printf("GetGameByIdController - error. %s", err)
			return
		}
	}
}

func (c *GamesController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var gdto dto.UpdateGameInfoRequestDto
		json.NewDecoder(r.Body).Decode(&gdto)
		gum := *mappers.ToGameUpdateModel(&gdto)
		err := c.gameService.UpdateById(id, gum)
		if err != nil {
			log.Printf("UpdateGameController - error on update. %s", err)
			return
		}
		game, err := c.gameService.GetById(id)
		if err != nil {
			log.Printf("UpdateGameController - error on update. %s", err)
			return
		}
		response := mappers.ToUpdateGameResponse(game)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}
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
		m := model.CreateMessageModel{
			ProjectId: id,
			Type:      dto.Type,
			Message:   dto.Message,
		}
		err = c.messageService.Send(id, &m)
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
