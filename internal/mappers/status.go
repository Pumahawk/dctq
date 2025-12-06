package mappers

import (
	"github.com/Pumahawk/dctq/internal/dto"
	"github.com/Pumahawk/dctq/internal/model"
)

func ToGetAllStatusResponseDto(status []model.StatusModel) *dto.GetAllStatusResponseDto {
	r := dto.GetAllStatusResponseDto{}
	for _, g := range status {
		gdto := dto.GetAllStatusResponseStatusDto{
			Id:   g.Id,
			Data: g.Data,
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
		Id:   status.Id,
		Data: status.Data,
	}
}

func ToGetStatusByIdResponseDto(status *model.StatusModel) *dto.GetStatusByIdResponseDto {
	if status == nil {
		return nil
	}
	return &dto.GetStatusByIdResponseDto{
		Id:   status.Id,
		Data: status.Data,
	}
}

func ToSimplStatusCreateInfoModel(status *dto.CreateStatusInfoRequestDto) *model.SimplStatusCreateInfoModel {
	if status == nil {
		return nil
	}
	return &model.SimplStatusCreateInfoModel{
		Data: status.Data,
	}
}

func ToUpdateStatusResponse(status *model.StatusModel) *dto.UpdateStatusResponseDto {
	if status == nil {
		return nil
	}
	return &dto.UpdateStatusResponseDto{
		Id:   status.Id,
		Data: status.Data,
	}
}

func ToStatusUpdateModel(status *dto.UpdateStatusInfoRequestDto) *model.StatusUpdateModel {
	if status == nil {
		return nil
	}
	return &model.StatusUpdateModel{
		Data: status.Data,
	}
}
