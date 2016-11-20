package main

import "fmt"

type Board struct {
	board [8][8]Mover
}

func (b *Board) initBoard() {
	//pawn loop
	for i := 0; i < 8; i++ {
		//white pawns
		var wtemp Pawn
		wtemp.GamePiece.Icon = WHITE_PAWN
		wtemp.GamePiece.IsWhite = true
		wtemp.GamePiece.Pos = Coord{x: i, y: 1}
		wtemp.IsFirst = true

		b.board[1][i] = wtemp

		//black pawns
		var btemp Pawn
		btemp.GamePiece.Icon = BLACK_PAWN
		btemp.GamePiece.IsWhite = true
		btemp.GamePiece.Pos = Coord{x: i, y: 6}
		btemp.IsFirst = true

		b.board[6][i] = btemp
	}

	//rooks
	b.initRooks()
	//knights
	b.initKnights()
	//bishops
	b.initBishops()
	//queen
	b.initQueens()
	//king
	b.initKings()
}
func (b *Board) initRooks() {
	//white
	b.board[0][0] = Rook{
		GamePiece{
			Icon:    WHITE_ROOK,
			IsWhite: true,
			Pos:     Coord{x: 0, y: 0},
		},
	}
	b.board[0][7] = Rook{
		GamePiece{
			Icon:    WHITE_ROOK,
			IsWhite: true,
			Pos:     Coord{x: 7, y: 0},
		},
	}
	//black
	b.board[7][0] = Rook{
		GamePiece{
			Icon:    BLACK_ROOK,
			IsWhite: false,
			Pos:     Coord{x: 0, y: 7},
		},
	}
	b.board[7][7] = Rook{
		GamePiece{
			Icon:    BLACK_ROOK,
			IsWhite: false,
			Pos:     Coord{x: 7, y: 7},
		},
	}
}
func (b *Board) initKnights() {
	//white
	b.board[0][1] = Knight{
		GamePiece{
			Icon:    WHITE_KNIGHT,
			IsWhite: true,
			Pos:     Coord{x: 1, y: 0},
		},
	}
	b.board[0][6] = Knight{
		GamePiece{
			Icon:    WHITE_KNIGHT,
			IsWhite: true,
			Pos:     Coord{x: 6, y: 0},
		},
	}
	//black
	b.board[7][1] = Knight{
		GamePiece{
			Icon:    BLACK_KNIGHT,
			IsWhite: false,
			Pos:     Coord{x: 1, y: 7},
		},
	}
	b.board[7][6] = Knight{
		GamePiece{
			Icon:    BLACK_KNIGHT,
			IsWhite: false,
			Pos:     Coord{x: 6, y: 7},
		},
	}
}
func (b *Board) initBishops() {
	//white
	b.board[0][2] = Bishop{
		GamePiece{
			Icon:    WHITE_BISHOP,
			IsWhite: true,
			Pos:     Coord{x: 2, y: 0},
		},
	}
	b.board[0][5] = Bishop{
		GamePiece{
			Icon:    WHITE_BISHOP,
			IsWhite: true,
			Pos:     Coord{x: 5, y: 0},
		},
	}
	//black
	b.board[7][2] = Bishop{
		GamePiece{
			Icon:    BLACK_BISHOP,
			IsWhite: false,
			Pos:     Coord{x: 2, y: 7},
		},
	}
	b.board[7][5] = Bishop{
		GamePiece{
			Icon:    BLACK_BISHOP,
			IsWhite: false,
			Pos:     Coord{x: 5, y: 7},
		},
	}
}
func (b *Board) initQueens() {
	//white
	b.board[0][3] = Queen{
		GamePiece{
			Icon:    WHITE_QUEEN,
			IsWhite: true,
			Pos:     Coord{x: 3, y: 0},
		},
	}

	//black
	b.board[7][3] = Queen{
		GamePiece{
			Icon:    BLACK_QUEEN,
			IsWhite: false,
			Pos:     Coord{x: 3, y: 7},
		},
	}
}
func (b *Board) initKings() {
	//white
	b.board[0][4] = King{
		GamePiece{
			Icon:    WHITE_KING,
			IsWhite: true,
			Pos:     Coord{x: 4, y: 0},
		},
	}

	//black
	b.board[7][4] = Queen{
		GamePiece{
			Icon:    BLACK_KING,
			IsWhite: false,
			Pos:     Coord{x: 4, y: 7},
		},
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
			var icon string
			if b.board[i][j] != nil {
				icon = b.board[i][j].getIcon()
			} else {
				icon = " "
			}
			fmt.Print(icon)
			fmt.Print("|")
		}
		fmt.Println("")
	}
	fmt.Println("+-----------------+")
}
