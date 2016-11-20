package main

import "fmt"

type Board struct {
	board [8][8]Mover
}

func (b *Board) initBoard() {
	for i := 0; i < 8; i++ {
		//white pawns
		var temp Pawn
		temp.GamePiece.Icon = WHITE_PAWN
		temp.GamePiece.IsWhite = true
		temp.GamePiece.Pos = Coord{x: i, y: 1}
		temp.IsFirst = true

		b.board[i][1] = temp
	}
}
func (b *Board) Print() {
	fmt.Println(" +----------------+")
	fmt.Println(" | 0 1 2 3 4 5 6 7 |")
	fmt.Println(" +----------------+")
	for i := 0; i < 8; i++ {
		fmt.Print(i)
		fmt.Print("|")
		for j := 0; j < 8; j++ {
			icon := b.board[i][j].getIcon()
			if icon == "" {
				icon = " "
			}
			fmt.Print(icon)
			fmt.Print("|")
		}
		fmt.Println("")
	}
	fmt.Println("+-----------------+")
}
