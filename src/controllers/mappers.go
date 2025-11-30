package controllers

import (
	"github.com/Pumahawk/cluedo/src/dto"
	"github.com/Pumahawk/cluedo/src/model"
)

func ToGetAllGameResponseDto(games []model.GameModel) (r dto.GetAllGameResponseDto) {
	for _, g := range games {
		gdto := dto.SimplGameInfoDto{
			Id:    g.Id,
			Label: g.Label,
		}
		r.Games = append(r.Games, gdto)
	}
	return
}

func ToCreateGameInfoResponseDto(game model.GameModel) dto.CreateGameInfoResponseDto {
	gdto := dto.CreateGameInfoResponseDto{
		Id:    game.Id,
		Label: game.Label,
	}
	return gdto
}

func ToGetGameByIdResponseDto(game *model.GameModel) dto.GetGameByIdResponseDto {
	gdto := dto.GetGameByIdResponseDto{
		Id:    game.Id,
		Label: game.Label,
	}
	return gdto
}

func toSimplGameCreateInfoModel(game dto.CreateGameInfoRequestDto) (r model.SimplGameCreateInfoModel) {
	r.Label = game.Label
	return
}
