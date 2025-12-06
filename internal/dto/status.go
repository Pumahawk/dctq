package dto

type StatusDataDto = map[string]any

type GetAllStatusResponseDto struct {
	Status []GetAllStatusResponseStatusDto `json:"status"`
}

type GetAllStatusResponseStatusDto struct {
	Id string `json:"id"`
}

type GetStatusByIdResponseDto struct {
	Id   string        `json:"id"`
	Data StatusDataDto `json:"data"`
}

type UpdateStatusInfoRequestDto struct {
	Data StatusDataDto `json:"data"`
}

type CreateStatusInfoResponseDto struct {
	Id   string        `json:"id"`
	Data StatusDataDto `json:"data"`
}

type CreateStatusInfoRequestDto struct {
	Data StatusDataDto `json:"data"`
}

type UpdateStatusResponseDto struct {
	Id   string        `json:"id"`
	Data StatusDataDto `json:"data"`
}
