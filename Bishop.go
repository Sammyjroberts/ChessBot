package main

type Bishop struct {
	GamePiece
}

func (r Bishop) getMoves(b *Board) []Coord {
	//if first turn, can go 2 forward
	var moves []Coord
	return moves
}
func (r Bishop) getIcon() string {
	return r.GamePiece.Icon
}
