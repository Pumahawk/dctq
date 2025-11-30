package services

import "github.com/Pumahawk/cluedo/src/model"

type GetAllGamesService = func() []model.GameModel
type CreateGameService = func(game model.SimplGameCreateInfoModel) []model.GameModel
type GetGameByIdService = func(id string) []model.GameModel
type UpdateGameByIdService = func(id string, info model.SimplGameUpdateInfoModel) model.GameModel
type SendMessageService = func(idGame string, info model.CreateMessageModel) model.MessageModel
