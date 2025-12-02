package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Pumahawk/cluedo/src/model"
)

var ErrGameNotFound = errors.New("Game not found")

type GameListener interface {
	GetAll() ([]model.GameModel, error)
}

type GameCreator interface {
	Create(model.SimplGameCreateInfoModel) (*model.GameModel, error)
}

type GameGetter interface {
	GetById(string) (*model.GameModel, error)
}

type GameUpdater interface {
	UpdateById(id string, game model.GameUpdateModel) error
}

type MessageSender interface {
	Send(string, *model.CreateMessageModel) (*model.MessageModel, error)
}

type GameServiceImpl struct {
	gameCounterId int64
	gamesInMemory model.ServerModel
}

type MessageServiceImpl struct {
}

func NewGameServiceImpl() *GameServiceImpl {
	return &GameServiceImpl{
		gamesInMemory: model.ServerModel{},
	}
}

func (s *GameServiceImpl) GetAll() ([]model.GameModel, error) {
	return s.gamesInMemory.Games, nil
}

func (s *GameServiceImpl) Create(game model.SimplGameCreateInfoModel) (*model.GameModel, error) {
	s.gameCounterId = s.gameCounterId + 1
	gameToCreate := model.GameModel{
		Id:    strconv.FormatInt(s.gameCounterId, 10),
		Label: game.Label,
	}
	s.gamesInMemory.Games = append(s.gamesInMemory.Games, gameToCreate)
	return &gameToCreate, nil
}

func (s *GameServiceImpl) GetById(id string) (*model.GameModel, error) {
	for i := range s.gamesInMemory.Games {
		if s.gamesInMemory.Games[i].Id == id {
			return &s.gamesInMemory.Games[i], nil
		}
	}
	return nil, ErrGameNotFound
}

func (s *GameServiceImpl) UpdateById(id string, game model.GameUpdateModel) error {
	for i, g := range s.gamesInMemory.Games {
		if id == g.Id {
			s.gamesInMemory.Games[i].Label = game.Label
			return nil
		}
	}
	return ErrGameNotFound
}

func (*MessageServiceImpl) Send(t string, message *model.CreateMessageModel) (*model.MessageModel, error) {
	return nil, fmt.Errorf("Message sender not implemented")
}
