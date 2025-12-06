package model

import "context"

type SimplGameInfoModel struct {
	Id    string
	Label string
}

type SimplGameCreateInfoModel struct {
	Label string
}

type GameUpdateModel struct {
	Label string
}

type GameModel struct {
	Id             string
	Label          string
	Players        []PlayerModel
	Cards          [3]CardModel
	MessageSockets []MessageSocket
}

type CardModel struct {
	Id    string
	Label string
}

type PlayerModel struct {
	Id    string
	Label string
	Cards []CardModel
	Pawn  PawnModel
}

type PawnModel struct {
	Color string
	Label string
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
	Games []GameModel
}
