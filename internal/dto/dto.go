package dto

type GetAllGameResponseDto struct {
	Games []SimplGameInfoDto
}

type SimplGameInfoDto struct {
	Id    string
	Label string
}

type GetGameByIdResponseDto struct {
	Id    string
	Label string
}

type UpdateGameInfoRequestDto struct {
	Label string
}

type CreateGameInfoResponseDto struct {
	Id    string
	Label string
}

type CreateGameInfoRequestDto struct {
	Label string
}

type UpdateGameResponseDto struct {
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
