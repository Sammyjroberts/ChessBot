package main

type Coord struct {
	x int
	y int
}
type GamePiece struct {
	Icon    string
	Pos     Coord
	IsWhite bool
}
