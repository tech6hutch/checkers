package board

import (
	"github.com/tech6hutch/checkers/model/pieces"
	"strconv"
	"strings"
	"unicode/utf8"
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

func (board *Board) String() string {
	var b strings.Builder

	for colN := 0; colN < Width; colN++ {
		n := string(colN)
		b.WriteString(n)
	}
	b.WriteRune('\n')

	for rowN, row := range board {
		// todo: instead, add to number first, to get letter
		rowLetter, _ := utf8.DecodeRuneInString(strconv.Itoa(rowN))
		// b.WriteRune(rowLetter)
		b.WriteString(rowLetter)
		for colN := range row {
			numStr := strconv.Itoa(colN)
			numRune, _ := utf8.DecodeRuneInString(numStr)
			b.WriteRune(numRune)
		}
	}
	for rowN, rowLetter := 0, 65; rowN < Width; rowN, rowLetter = rowN+1, rowLetter+1 {
		row := board[rowN]
		for colN, p := range row {

		}
	}
	for rowN, row := range board {
		rowLetter := rune(rowN + 65)
		b.WriteRune(rowLetter)
		for _, p := range row {
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
		b.WriteRune('\n')
	}
	return b.String()
}
