package main

type King struct {
	GamePiece
}

func (r King) getMoves(b *Board) []Coord {
	//if first turn, can go 2 forward
	var moves []Coord
	return moves
}
func (r King) getIcon() string {
	return r.GamePiece.Icon
}
