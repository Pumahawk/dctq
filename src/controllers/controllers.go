package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Pumahawk/cluedo/src/dto"
	"github.com/Pumahawk/cluedo/src/services"
)

var GamesEndpoint = "/games"
var GameByIdEndpoint = "/games/:id"
var MessagesEndpoint = "/games/:id/messages"

func GamesController(getAllGamesService services.GetAllGamesService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		games := getAllGamesService()
		gamesDto := ToGetAllGameResponseDto(games)
		w.WriteHeader(200)
		err := json.NewEncoder(w).Encode(gamesDto)
		if err != nil {
			log.Printf("GamesController - error. %s", err)
		}
	}
}

func CreateGameController(createGameService services.CreateGameService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createGameInfoRequestDto dto.CreateGameInfoRequestDto
		json.NewDecoder(r.Body).Decode(&createGameInfoRequestDto)
		game := createGameService(toSimplGameCreateInfoModel(createGameInfoRequestDto))
		responseDto := ToCreateGameInfoResponseDto(game)
		w.WriteHeader(200)
		err := json.NewEncoder(w).Encode(responseDto)
		if err != nil {
			log.Printf("CreateGameController - error. %s", err)
		}
	}
}

func GetGameByIdController(getGameByIdService services.GetGameByIdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		gameModel := getGameByIdService(id)
		responseDto := ToGetGameByIdResponseDto(gameModel)
		w.WriteHeader(200)
		err := json.NewEncoder(w).Encode(responseDto)
		if err != nil {
			log.Printf("GetGameByIdController - error. %s", err)
		}
	}
}

func UpdateGetGameByIdController(getGameByIdService services.GetGameByIdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("API GameByIdController - Not yet implemented")
		w.WriteHeader(500)
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
