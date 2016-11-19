package main

import "fmt"

type Board struct {
	board [8][8]GamePiece
}

func (b *Board) initBoard() {
	b.board = [8][8]GamePiece{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
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
			icon := b.board[i][j].Icon
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
