package mappers

import (
	"github.com/Pumahawk/dctq/internal/dto"
	"github.com/Pumahawk/dctq/internal/model"
)

func ToGetAllStatusResponseDto(status []model.StatusModel) *dto.GetAllStatusResponseDto {
	r := dto.GetAllStatusResponseDto{}
	for _, g := range status {
		gdto := dto.SimplStatusInfoDto{
			Id:    g.Id,
			Label: g.Label,
		}
		r.Status = append(r.Status, gdto)
	}
	return &r
}

func ToCreateStatusInfoResponseDto(status *model.StatusModel) *dto.CreateStatusInfoResponseDto {
	if status == nil {
		return nil
	}
	return &dto.CreateStatusInfoResponseDto{
		Id:    status.Id,
		Label: status.Label,
	}
}

func ToGetStatusByIdResponseDto(status *model.StatusModel) *dto.GetStatusByIdResponseDto {
	if status == nil {
		return nil
	}
	return &dto.GetStatusByIdResponseDto{
		Id:    status.Id,
		Label: status.Label,
	}
}

func ToSimplStatusCreateInfoModel(status *dto.CreateStatusInfoRequestDto) *model.SimplStatusCreateInfoModel {
	if status == nil {
		return nil
	}
	return &model.SimplStatusCreateInfoModel{
		Label: status.Label,
	}
}

func ToUpdateStatusResponse(status *model.StatusModel) *dto.UpdateStatusResponseDto {
	if status == nil {
		return nil
	}
	return &dto.UpdateStatusResponseDto{
		Id:    status.Id,
		Label: status.Label,
	}
}

func ToStatusUpdateModel(status *dto.UpdateStatusInfoRequestDto) *model.StatusUpdateModel {
	if status == nil {
		return nil
	}
	return &model.StatusUpdateModel{
		Label: status.Label,
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
