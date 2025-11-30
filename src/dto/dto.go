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
