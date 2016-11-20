package main

type Knight struct {
	GamePiece
}

func (r Knight) getMoves(b *Board) []Coord {
	//if first turn, can go 2 forward
	var moves []Coord
	return moves
}
func (r Knight) getIcon() string {
	return r.GamePiece.Icon
}
