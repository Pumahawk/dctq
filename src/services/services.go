package services

import "github.com/Pumahawk/cluedo/src/model"

type GetAllGamesService = func() ([]model.GameModel, error)
type CreateGameService = func(game *model.SimplGameCreateInfoModel) (*model.GameModel, error)
type GetGameByIdService = func(id string) (*model.GameModel, error)
type UpdateGameService = func(game *model.GameModel) error
type SendMessageService = func(idGame string, info *model.CreateMessageModel) (*model.MessageModel, error)
