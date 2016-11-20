package main

type Coord struct {
	x int
	y int
}
type GamePiece struct {
	Icon string
	//redundant but makes things way easier to code
	Pos     Coord
	IsWhite bool
}

type Mover interface {
	getMoves(b *Board) []Coord
	getIcon() string
}
