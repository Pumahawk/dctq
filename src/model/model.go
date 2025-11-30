package model

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
