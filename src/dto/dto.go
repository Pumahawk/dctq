package dto

type GetAllGameResponseDto struct {
	games []SimplGameInfoDto
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
