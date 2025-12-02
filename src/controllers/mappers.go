package controllers

import (
	"github.com/Pumahawk/cluedo/src/dto"
	"github.com/Pumahawk/cluedo/src/model"
)

func ToGetAllGameResponseDto(games []model.GameModel) *dto.GetAllGameResponseDto {
	r := dto.GetAllGameResponseDto{}
	for _, g := range games {
		gdto := dto.SimplGameInfoDto{
			Id:    g.Id,
			Label: g.Label,
		}
		r.Games = append(r.Games, gdto)
	}
	return &r
}

func ToCreateGameInfoResponseDto(game *model.GameModel) *dto.CreateGameInfoResponseDto {
	if game == nil {
		return nil
	}
	gdto := dto.CreateGameInfoResponseDto{
		Id:    game.Id,
		Label: game.Label,
	}
	return &gdto
}

func ToGetGameByIdResponseDto(game *model.GameModel) *dto.GetGameByIdResponseDto {
	if game == nil {
		return nil
	}
	gdto := dto.GetGameByIdResponseDto{
		Id:    game.Id,
		Label: game.Label,
	}
	return &gdto
}

func toSimplGameCreateInfoModel(game *dto.CreateGameInfoRequestDto) *model.SimplGameCreateInfoModel {
	if game == nil {
		return nil
	}
	r := model.SimplGameCreateInfoModel{}
	r.Label = game.Label
	return &r
}

func toUpdateGameResponse(game *model.GameModel) *dto.UpdateGameResponseDto {
	if game == nil {
		return nil
	}
	dto := dto.UpdateGameResponseDto{
		Id:    game.Id,
		Label: game.Label,
	}
	return &dto
}

func toGameUpdateModel(game *dto.UpdateGameInfoRequestDto) *model.GameUpdateModel {
	if game == nil {
		return nil
	}
	m := model.GameUpdateModel{
		Label: game.Label,
	}
	return &m
}
