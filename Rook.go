package main

type Rook struct {
	GamePiece
}

func (r Rook) getMoves(b *Board) []Coord {
	//if first turn, can go 2 forward
	var moves []Coord
	return moves
}
func (r Rook) getIcon() string {
	return r.GamePiece.Icon
}
