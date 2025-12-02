package model

type SimplGameInfoModel struct {
	Id    string
	Label string
}

type SimplGameCreateInfoModel struct {
	Label string
}

type GameUpdateModel struct {
	Id      string
	Label   string
	Players []PlayerModel
	Cards   [3]CardModel
}

type GameModel struct {
	Id      string
	Label   string
	Players []PlayerModel
	Cards   [3]CardModel
}

type CardModel struct {
	Id    string
	Label string
}

type PlayerModel struct {
	Id    string
	Label string
	Cards []CardModel
	Pawn  *PawnModel
}

type PawnModel struct {
	Color string
	Label string
}

type MessageModel struct {
	Id      string
	Time    string
	Type    string
	Message map[string]string
}

type CreateMessageModel struct {
	Type    string
	Message map[string]string
}

type ServerModel struct {
	Games []GameModel
}
