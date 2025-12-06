package mappers

import (
	"github.com/Pumahawk/dctq/internal/dto"
	"github.com/Pumahawk/dctq/internal/model"
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
	return &dto.CreateGameInfoResponseDto{
		Id:    game.Id,
		Label: game.Label,
	}
}

func ToGetGameByIdResponseDto(game *model.GameModel) *dto.GetGameByIdResponseDto {
	if game == nil {
		return nil
	}
	return &dto.GetGameByIdResponseDto{
		Id:    game.Id,
		Label: game.Label,
	}
}

func ToSimplGameCreateInfoModel(game *dto.CreateGameInfoRequestDto) *model.SimplGameCreateInfoModel {
	if game == nil {
		return nil
	}
	return &model.SimplGameCreateInfoModel{
		Label: game.Label,
	}
}

func ToUpdateGameResponse(game *model.GameModel) *dto.UpdateGameResponseDto {
	if game == nil {
		return nil
	}
	return &dto.UpdateGameResponseDto{
		Id:    game.Id,
		Label: game.Label,
	}
}

func ToGameUpdateModel(game *dto.UpdateGameInfoRequestDto) *model.GameUpdateModel {
	if game == nil {
		return nil
	}
	return &model.GameUpdateModel{
		Label: game.Label,
	}
}
func ToCreateMessageModelFromDto(id string, dto *dto.SendMessageRequestDto) *model.CreateMessageModel {
	if dto == nil {
		return nil
	}
	return &model.CreateMessageModel{
		ProjectId: id,
		Type:      dto.Type,
		Message:   dto.Message,
	}
}
