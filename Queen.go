package main

type Queen struct {
	GamePiece
}

func (r Queen) getMoves(b *Board) []Coord {
	//if first turn, can go 2 forward
	var moves []Coord
	return moves
}
func (r Queen) getIcon() string {
	return r.GamePiece.Icon
}
