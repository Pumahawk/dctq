package dto

type MessageDataDto = map[string]any

type SendMessageRequestDto struct {
	Type    string
	Message MessageDataDto
}

type CreateMessageResponseDto struct {
	Type    string
	Message MessageDataDto
}

type FollowMessageResponseDto struct {
	Type    string
	Message MessageDataDto
}
