package services

import "github.com/Pumahawk/cluedo/src/model"

func GetAllGamesServiceImpl() ([]model.GameModel, error) {
	return gamesInMemory.Games, nil
}
