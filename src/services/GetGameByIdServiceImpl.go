package services

import (
	"github.com/Pumahawk/cluedo/src/model"
)

func GetGameByIdServiceImpl(id string) *model.GameModel {
	for _, g := range gamesInMemory.Games {
		if g.Id == id {
			return &g
		}
	}
	return nil
}
