package services

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/Pumahawk/dctq/internal/model"
)

var ErrStatusNotFound = errors.New("Status not found")

type StatusService interface {
	StatusListener
	StatusCreator
	StatusGetter
	StatusUpdater
	StatusMessageFollow
}

type MessageService interface {
	MessageSender
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
	Send(*model.CreateMessageModel) error
}

type StatusMessageFollow interface {
	FollowMessages(c context.Context, projectId string) (<-chan model.MessageModel, error)
}

type ServerMessageProcessor interface {
	Start() error
}

type ServerMessageProcessorImpl struct {
	statusService        StatusService
	ctx                  context.Context
	globalMessageChannel chan model.CreateMessageModel
}

func NewServerMessageProcessorImpl(ctx context.Context, globalMessageChannel chan model.CreateMessageModel, statusService StatusService) *ServerMessageProcessorImpl {
	return &ServerMessageProcessorImpl{
		ctx:                  ctx,
		globalMessageChannel: globalMessageChannel,
		statusService:        statusService,
	}
}

type StatusServiceImpl struct {
	statusInMemory model.ServerModel
}

type MessageServiceImpl struct {
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
	id, _ := GenerateUUIDv4()
	statusToCreate := model.StatusModel{
		Id:   id,
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

func (s *StatusServiceImpl) FollowMessages(context context.Context, projectId string) (<-chan model.MessageModel, error) {
	status, err := s.GetById(projectId)
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

func NewMessageServiceImpl(globalMessageChannel chan model.CreateMessageModel) *MessageServiceImpl {
	return &MessageServiceImpl{
		globalMessageChannel: globalMessageChannel,
	}
}

func (m *MessageServiceImpl) Send(message *model.CreateMessageModel) error {
	m.globalMessageChannel <- *message
	return nil
}

func (s *ServerMessageProcessorImpl) Start() error {
	log.Printf("MessageServiceImpl - Start message processor")
	for {
		select {
		case <-s.ctx.Done():
			log.Printf("MessageServiceImpl - Close server")
			// TODO
		case message := <-s.globalMessageChannel:
			log.Printf("MessageServiceImpl - Incoming message")
			statusId := message.ProjectId
			status, err := s.statusService.GetById(statusId)
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

func GenerateUUIDv4() (string, error) {
	uuid := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, uuid)
	if err != nil {
		return "", err
	}
	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4],
		uuid[4:6],
		uuid[6:8],
		uuid[8:10],
		uuid[10:16],
	), nil
}
