package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pumahawk/dctq/internal/dto"
	"github.com/Pumahawk/dctq/internal/mappers"
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

type MessagesController struct {
}

func NewGamesController(gameService services.GameService) *GamesController {
	return &GamesController{
		gameService,
	}
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
		log.Println("API SandMessageController - Not yet implemented")
		w.WriteHeader(500)
	}
}

func (c *MessagesController) Follow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("API FollowMessageController - Not yet implemented")
		w.WriteHeader(500)
	}
}
