package services

import (
	"github.com/Pumahawk/cluedo/src/model"
)

func GetGameByIdServiceImpl(id string) (*model.GameModel, error) {
	for _, g := range gamesInMemory.Games {
		if g.Id == id {
			return &g, nil
		}
	}
	return nil, nil
}
