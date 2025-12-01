package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pumahawk/cluedo/src/dto"
	"github.com/Pumahawk/cluedo/src/services"
)

var GamesEndpoint = "/games"
var GameByIdEndpoint = "/games/{id}"
var MessagesEndpoint = "/games/{id}/messages"

func GamesController(getAllGamesService services.GetAllGamesService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		games, err := getAllGamesService()
		if err != nil {
			log.Printf("GamesController - error. %s", err)
			return
		}
		gamesDto := ToGetAllGameResponseDto(games)
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(gamesDto)
		if err != nil {
			log.Printf("GamesController - error. %s", err)
		}
	}
}

func CreateGameController(createGameService services.CreateGameService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createGameInfoRequestDto dto.CreateGameInfoRequestDto
		json.NewDecoder(r.Body).Decode(&createGameInfoRequestDto)
		game, err := createGameService(toSimplGameCreateInfoModel(&createGameInfoRequestDto))
		if err != nil {
			log.Printf("CreateGameController - error. %s", err)
			return
		}
		responseDto := ToCreateGameInfoResponseDto(game)
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(responseDto)
		if err != nil {
			log.Printf("CreateGameController - error. %s", err)
			return
		}
	}
}

func GetGameByIdController(getGameByIdService services.GetGameByIdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		gameModel, err := getGameByIdService(id)
		if err != nil {
			log.Printf("GetGameByIdController - error. %s", err)
			return
		}
		responseDto := ToGetGameByIdResponseDto(gameModel)
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(responseDto)
		if err != nil {
			log.Printf("GetGameByIdController - error. %s", err)
			return
		}
	}
}

func UpdateGameController(getGameByIdService services.GetGameByIdService, updateGameService services.UpdateGameService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		game, err := getGameByIdService(id)
		if err != nil {
			log.Printf("UpdateGetGameByIdController - error on get game. %s", err)
			return
		}
		if game != nil {
			log.Printf("UpdateGetGameByIdController - not found. %s", err)
			return
		}
		err = updateGameService(game)
		if err != nil {
			log.Printf("UpdateGameController - error on update. %s", err)
			return
		}
		response := toUpdateGameResponse(game)
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}
}

func SandMessageController(sendMessageService services.SendMessageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("API SandMessageController - Not yet implemented")
		w.WriteHeader(500)
	}
}

func FollowMessageController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("API FollowMessageController - Not yet implemented")
		w.WriteHeader(500)
	}
}
