package services

import (
	"github.com/Pumahawk/cluedo/src/model"
	"github.com/google/uuid"
)

func CreateGameServiceImpl(game model.SimplGameCreateInfoModel) model.GameModel {
	gameToCreate := model.GameModel{
		Id:    uuid.New().String(),
		Label: game.Label,
	}
	gamesInMemory.Games = append(gamesInMemory.Games, gameToCreate)
	return gameToCreate
}
