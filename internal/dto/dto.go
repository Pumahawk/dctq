package dto

type GetAllStatusResponseDto struct {
	Status []SimplStatusInfoDto
}

type SimplStatusInfoDto struct {
	Id    string
	Label string
}

type GetStatusByIdResponseDto struct {
	Id    string
	Label string
}

type UpdateStatusInfoRequestDto struct {
	Label string
}

type CreateStatusInfoResponseDto struct {
	Id    string
	Label string
}

type CreateStatusInfoRequestDto struct {
	Label string
}

type UpdateStatusResponseDto struct {
	Id    string
	Label string
}

type SendMessageRequestDto struct {
	Type    string
	Message map[string]string
}

type CreateMessageResponseDto struct {
	Type    string
	Message map[string]string
}

type FollowMessageResponseDto struct {
	Type    string
	Message map[string]string
}
