package board

import (
	"github.com/tech6hutch/checkers/model/pieces"
	"strings"
)

const Width = 8

type Board [Width][Width]*pieces.Piece

func New() *Board {
	board := new(Board)
	for row := 0; row < 3; row++ {
		var col int
		if row%2 == 0 {
			col = 1
		} else {
			col = 0
		}
		for ; col < Width; col += 2 {
			p := pieces.NewRed()
			board[row][col] = p
		}
	}
	for row := Width - 1; row > Width-4; row-- {
		var col int
		if row%2 == 0 {
			col = 1
		} else {
			col = 0
		}
		for ; col < Width; col += 2 {
			p := pieces.NewBlack()
			board[row][col] = p
		}
	}
	return board
}

func (board *Board) String() string {
	var b strings.Builder
	for row := range board {
		for col := range board[row] {
			p := board[row][col]
			if p == nil {
				b.WriteRune('#')
			} else {
				b.WriteRune('*')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}
