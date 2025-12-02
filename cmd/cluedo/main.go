package main

import (
	"log"
	"net/http"

	"github.com/Pumahawk/cluedo/internal/controllers"
	"github.com/Pumahawk/cluedo/internal/services"
)

func main() {
	log.Println("Starting cluedo server.")
	gameService := services.NewGameServiceImpl()
	gameController := controllers.NewGamesController(gameService)
	http.Handle("GET "+controllers.GamesEndpoint, gameController.GetAll())
	http.Handle("POST "+controllers.GamesEndpoint, gameController.Create())
	http.Handle("GET "+controllers.GameByIdEndpoint, gameController.GetById())
	http.Handle("PUT "+controllers.GameByIdEndpoint, gameController.Update())
	log.Println("Start Cluedo server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
