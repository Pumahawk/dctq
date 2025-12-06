package dto

type StatusDataDto = map[string]any

type GetAllStatusResponseDto struct {
	Status []GetAllStatusResponseStatusDto
}

type GetAllStatusResponseStatusDto struct {
	Id string
}

type GetStatusByIdResponseDto struct {
	Id   string
	Data StatusDataDto
}

type UpdateStatusInfoRequestDto struct {
	Data StatusDataDto
}

type CreateStatusInfoResponseDto struct {
	Id   string
	Data StatusDataDto
}

type CreateStatusInfoRequestDto struct {
	Data StatusDataDto
}

type UpdateStatusResponseDto struct {
	Id   string
	Data StatusDataDto
}
