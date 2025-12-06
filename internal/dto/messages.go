package dto

type MessageDataDto = map[string]any

type SendMessageRequestDto struct {
	Type    string         `json:"type"`
	Message MessageDataDto `json:"message"`
}

type CreateMessageResponseDto struct {
	Type    string         `json:"type"`
	Message MessageDataDto `json:"message"`
}

type FollowMessageResponseDto struct {
	Type    string         `json:"type"`
	Message MessageDataDto `json:"message"`
}
