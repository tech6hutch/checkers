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
		board.initRow(rowN, pieces.ColorLight)
	}
	for rowN := Width - 1; rowN > Width-4; rowN-- {
		board.initRow(rowN, pieces.ColorDark)
	}
	board.verifyPiecePositions()
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

func (board *Board) verifyPiecePositions() {
	for rowN, row := range board {
		for colN, p := range row {
			if rowN%2 == colN%2 && p != nil {
				panic("wrong piece placement")
			}
		}
	}
}

func (board Board) String() string {
	var b strings.Builder

	b.WriteRune(' ')
	for colN := range board[0] {
		colLetter := rune(colN%26 + 'a')
		b.WriteRune(colLetter)
	}
	columnLetters := b.String()
	b.WriteRune('\n')

	for rowN := Width-1; rowN >= 0; rowN-- {
		rowDigit := rune((rowN+1)%10 + '0')
		b.WriteRune(rowDigit)

		for _, p := range board[rowN] {
			switch {
			case p == nil:
				b.WriteRune('#')
			case p.Kind == pieces.Man:
				b.WriteRune('*')
			case p.Kind == pieces.King:
				b.WriteRune('^')
			default:
				panic("unknown piece kind")
			}
		}

		b.WriteRune(rowDigit)
		b.WriteRune('\n')
	}

	b.WriteString(columnLetters)

	return b.String()
}
