package model

type Game struct {
	Id      string
	Label   string
	Players []Player
	Cards   [3]Card
}

type Card struct {
	Id    string
	Label string
}

type Player struct {
	Id    string
	Label string
	Cards []Card
	Pawn  *Pawn
}

type Pawn struct {
	Color string
	Label string
}
