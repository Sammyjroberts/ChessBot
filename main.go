package main

type Pawn struct {
	GamePiece
	IsFirst bool
}
type Mover interface {
	getMoves(b *Board) []Coord
}

func (p Pawn) getMoves(b *Board) {
	//if first turn, can go 2 forward
	var moves []Coord
	yDiff := 1
	if p.IsFirst {
		if !p.GamePiece.isWhite {
			yDiff = -1
		}
		if p.pos.y+(yDiff*2) < 8 {
			moves = append(moves, Coord{y: p.pos.y + 2, x: p.pos.x})
		}

	}
	//if can en passe
	//else can
}

func main() {
	var b Board
	b.Print()
}
