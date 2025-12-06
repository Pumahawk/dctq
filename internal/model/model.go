package model

import "context"

type StatusDataModel = map[string]any
type MessageDataModel = map[string]any

type SimplStatusInfoModel struct {
	Id string
}

type SimplStatusCreateInfoModel struct {
	Data StatusDataModel
}

type StatusUpdateModel struct {
	Data StatusDataModel
}

type StatusModel struct {
	Id             string
	Data           StatusDataModel
	MessageSockets []MessageSocket
}

type MessageSocket struct {
	Context context.Context
	Channel chan MessageModel
}

type MessageModel struct {
	Type    string
	Message MessageDataModel
}

type CreateMessageModel struct {
	ProjectId string
	Type      string
	Message   MessageDataModel
}

type ServerModel struct {
	Status []StatusModel
}
