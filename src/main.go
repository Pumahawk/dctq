package main

import (
	"log"
	"net/http"

	"github.com/Pumahawk/cluedo/src/controllers"
	"github.com/Pumahawk/cluedo/src/services"
)

func main() {
	log.Println("Starting cluedo server.")
	http.Handle("GET "+controllers.GamesEndpoint, controllers.GamesController(services.GetAllGamesServiceImpl))
	http.Handle("POST "+controllers.GamesEndpoint, controllers.CreateGameController(services.CreateGameServiceImpl))
	http.Handle("GET "+controllers.GameByIdEndpoint, controllers.GetGameByIdController(services.GetGameByIdServiceImpl))
	log.Println("Start Cluedo server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
