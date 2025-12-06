package model

import "context"

type StatusDataModel = map[string]any

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
	Message map[string]string
}

type CreateMessageModel struct {
	ProjectId string
	Type      string
	Message   map[string]string
}

type ServerModel struct {
	Status []StatusModel
}
