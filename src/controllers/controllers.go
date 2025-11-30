package controllers

import (
	"log"
	"net/http"

	"github.com/Pumahawk/cluedo/src/services"
)

var GamesEndpoint = "/games"
var GameByIdEndpoint = "/games/{id}"
var MessagesEndpoint = "/games/{id}/messages"

func GamesController(getAllGamesService services.GetAllGamesService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("API GamesController - Not yet implemented")
		w.WriteHeader(500)
	}
}

func CreateGameController(createGameService services.CreateGameService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("API CreateGameController - Not yet implemented")
		w.WriteHeader(500)
	}
}

func GetGameByIdController(getGameByIdService services.GetGameByIdService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("API GameByIdController - Not yet implemented")
		w.WriteHeader(500)
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
