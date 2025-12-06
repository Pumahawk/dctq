package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/Pumahawk/dctq/internal/model"
)

var ErrStatusNotFound = errors.New("Status not found")

type StatusService interface {
	StatusListener
	StatusCreator
	StatusGetter
	StatusUpdater
}

type MessageService interface {
	MessageSender
	MessageFollow
}

type StatusListener interface {
	GetAll() ([]model.StatusModel, error)
}

type StatusCreator interface {
	Create(model.SimplStatusCreateInfoModel) (*model.StatusModel, error)
}

type StatusGetter interface {
	GetById(string) (*model.StatusModel, error)
}

type StatusUpdater interface {
	UpdateById(id string, status model.StatusUpdateModel) error
}

type MessageSender interface {
	Send(string, *model.CreateMessageModel) error
}

type MessageFollow interface {
	Follow(c context.Context, projectId string) (<-chan model.MessageModel, error)
}

type StatusServiceImpl struct {
	statusCounterId int64
	statusInMemory  model.ServerModel
}

type MessageServiceImpl struct {
	statusService        StatusService
	serverContext        context.Context
	globalMessageChannel chan model.CreateMessageModel
}

func NewStatusServiceImpl() *StatusServiceImpl {
	return &StatusServiceImpl{
		statusInMemory: model.ServerModel{},
	}
}

func (s *StatusServiceImpl) GetAll() ([]model.StatusModel, error) {
	return s.statusInMemory.Status, nil
}

func (s *StatusServiceImpl) Create(status model.SimplStatusCreateInfoModel) (*model.StatusModel, error) {
	s.statusCounterId = s.statusCounterId + 1
	statusToCreate := model.StatusModel{
		Id:   strconv.FormatInt(s.statusCounterId, 10),
		Data: status.Data,
	}
	s.statusInMemory.Status = append(s.statusInMemory.Status, statusToCreate)
	return &statusToCreate, nil
}

func (s *StatusServiceImpl) GetById(id string) (*model.StatusModel, error) {
	for i := range s.statusInMemory.Status {
		if s.statusInMemory.Status[i].Id == id {
			return &s.statusInMemory.Status[i], nil
		}
	}
	return nil, ErrStatusNotFound
}

func (s *StatusServiceImpl) UpdateById(id string, status model.StatusUpdateModel) error {
	for i, g := range s.statusInMemory.Status {
		if id == g.Id {
			s.statusInMemory.Status[i].Data = status.Data
			return nil
		}
	}
	return ErrStatusNotFound
}

func NewMessageServiceImpl(statusService StatusService) *MessageServiceImpl {
	messageServiceImpl := MessageServiceImpl{
		statusService:        statusService,
		serverContext:        context.TODO(),
		globalMessageChannel: make(chan model.CreateMessageModel),
	}
	return &messageServiceImpl
}

func (m *MessageServiceImpl) Send(projectId string, message *model.CreateMessageModel) error {
	_, err := m.statusService.GetById(projectId)
	if err != nil && err != ErrStatusNotFound {
		return fmt.Errorf("StatusService Send - unable to retrieve status by id %s", projectId)
	}
	if err == ErrStatusNotFound {
		return ErrStatusNotFound
	}
	m.globalMessageChannel <- *message
	return nil
}

func (m *MessageServiceImpl) Follow(context context.Context, projectId string) (<-chan model.MessageModel, error) {
	status, err := m.statusService.GetById(projectId)
	if err != nil && err != ErrStatusNotFound {
		return nil, fmt.Errorf("StatusService Follow - unable to retrieve status by id %s", projectId)
	}
	if err == ErrStatusNotFound {
		return nil, ErrStatusNotFound
	}
	channel := make(chan model.MessageModel)
	status.MessageSockets = append(status.MessageSockets, model.MessageSocket{
		Context: context,
		Channel: channel,
	})
	return channel, nil
}

func (m *MessageServiceImpl) StartServerMessageProcessor() error {
	log.Printf("MessageServiceImpl - Start message processor")
	for {
		select {
		case <-m.serverContext.Done():
			log.Printf("MessageServiceImpl - Close server")
			// TODO
		case message := <-m.globalMessageChannel:
			log.Printf("MessageServiceImpl - Incoming message")
			statusId := message.ProjectId
			status, err := m.statusService.GetById(statusId)
			if err != nil {
				log.Printf("MessageServiceImpl - Unable to retrieve project. %s", err)
				continue
			}
			for i := range status.MessageSockets {
				select {
				case <-status.MessageSockets[i].Context.Done():
				case status.MessageSockets[i].Channel <- model.MessageModel{
					// TODO
					Type:    message.Type,
					Message: message.Message,
				}:
				}
			}
		}
	}
}
