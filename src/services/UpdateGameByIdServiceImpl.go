package services

import (
	"fmt"
	"log"

	"github.com/Pumahawk/cluedo/src/model"
)

func UpdateGameServiceImpl(game *model.GameModel) error {
	gs := gamesInMemory.Games
	for i := range gs {
		if game == &gs[i] {
			log.Printf("UpdateGameServiceImpl - element updated")
			return nil
		}
	}
	return fmt.Errorf("element not found")
}
