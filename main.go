package main

type Pawn struct {
	GamePiece
	IsFirst bool
}
type Mover interface {
	getMoves(b *Board) []Coord
	getIcon() string
}

func (p Pawn) getMoves(b *Board) []Coord {
	//if first turn, can go 2 forward
	var moves []Coord
	yDiff := 1
	if p.IsFirst {
		if !p.GamePiece.IsWhite {
			yDiff = -1
		}
		if p.GamePiece.Pos.y+(yDiff*2) < 8 {
			moves = append(moves, Coord{y: p.GamePiece.Pos.y + 2, x: p.GamePiece.Pos.x})
		}

	}
	//if can en passe
	//else can
	return moves
}
func (p Pawn) getIcon() string {
	return p.GamePiece.Icon
}

func main() {
	var b Board
	b.initBoard()
	b.Print()
}
