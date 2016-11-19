package main

type Coord struct {
	x int
	y int
}
type GamePiece struct {
	Icon    string
	pos     Coord
	isWhite bool
}
