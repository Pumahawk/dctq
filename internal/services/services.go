package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/Pumahawk/dctq/internal/model"
)

var ErrGameNotFound = errors.New("Game not found")

type GameService interface {
	GameListener
	GameCreator
	GameGetter
	GameUpdater
}

type MessageService interface {
	MessageSender
	MessageFollow
}

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
	Send(string, *model.CreateMessageModel) error
}

type MessageFollow interface {
	Follow(c context.Context, projectId string) (<-chan model.MessageModel, error)
}

type GameServiceImpl struct {
	gameCounterId int64
	gamesInMemory model.ServerModel
}

type MessageServiceImpl struct {
	gameService          GameService
	serverContext        context.Context
	globalMessageChannel chan model.CreateMessageModel
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

func NewMessageServiceImpl(gameService GameService) *MessageServiceImpl {
	messageServiceImpl := MessageServiceImpl{
		gameService:          gameService,
		serverContext:        context.TODO(),
		globalMessageChannel: make(chan model.CreateMessageModel),
	}
	return &messageServiceImpl
}

func (m *MessageServiceImpl) Send(projectId string, message *model.CreateMessageModel) error {
	_, err := m.gameService.GetById(projectId)
	if err != nil && err != ErrGameNotFound {
		return fmt.Errorf("GameService Send - unable to retrieve game by id %s", projectId)
	}
	if err == ErrGameNotFound {
		return ErrGameNotFound
	}
	m.globalMessageChannel <- *message
	return nil
}

func (m *MessageServiceImpl) Follow(context context.Context, projectId string) (<-chan model.MessageModel, error) {
	game, err := m.gameService.GetById(projectId)
	if err != nil && err != ErrGameNotFound {
		return nil, fmt.Errorf("GameService Follow - unable to retrieve game by id %s", projectId)
	}
	if err == ErrGameNotFound {
		return nil, ErrGameNotFound
	}
	channel := make(chan model.MessageModel)
	game.MessageSockets = append(game.MessageSockets, model.MessageSocket{
		Context: context,
		Channel: channel,
	})
	return channel, nil
}

func (m *MessageServiceImpl) StartServerMessageProcessor() error {
	log.Printf("Start message processor")
	for {
		select {
		case <-m.serverContext.Done():
			log.Printf("Close server")
			// TODO
		case message := <-m.globalMessageChannel:
			log.Printf("Incoming message")
			gameId := message.ProjectId
			game, err := m.gameService.GetById(gameId)
			if err != nil {
				log.Printf("Unable to retrieve project. %s", err)
				continue
			}
			for i := range game.MessageSockets {
				select {
				case <-game.MessageSockets[i].Context.Done():
				case game.MessageSockets[i].Channel <- model.MessageModel{
					// TODO
					Type:    message.Type,
					Message: message.Message,
				}:
				}
			}
		}
	}
}
