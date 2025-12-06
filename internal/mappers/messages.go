package mappers

import (
	"github.com/Pumahawk/dctq/internal/dto"
	"github.com/Pumahawk/dctq/internal/model"
)

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
