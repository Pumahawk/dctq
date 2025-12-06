package dto

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
