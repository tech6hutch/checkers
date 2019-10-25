package board

import (
	"github.com/tech6hutch/checkers/model/pieces"
	"strings"
)

const Width = 8

type Board [Width][Width]*pieces.Piece

func New() *Board {
	board := new(Board)
	for rowN := 0; rowN < 3; rowN++ {
		board.initRow(rowN, pieces.ColorDark)
	}
	for rowN := Width - 1; rowN > Width-4; rowN-- {
		board.initRow(rowN, pieces.ColorLight)
	}
	// verify
	for rowN, row := range board {
		for colN, p := range row {
			if rowN % 2 == colN % 2 && p != nil {
				panic("wrong piece placement")
			}
		}
	}
	return board
}

func (board *Board) initRow(rowN int, color pieces.Color) {
	var colN int
	if rowN%2 == 0 {
		colN = 1
	} else {
		colN = 0
	}
	for ; colN < Width; colN += 2 {
		p := pieces.New(color)
		board[rowN][colN] = p
	}
}

func (board *Board) String() string {
	var b strings.Builder
	for _, row := range board {
		for _, p := range row {
			switch {
			case p == nil:              b.WriteRune('#')
			case p.Kind == pieces.Man:  b.WriteRune('*')
			case p.Kind == pieces.King: b.WriteRune('^')
			default: panic("unknown piece kind")
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}
