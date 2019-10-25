package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Checkers!")
	board := NewBoard()

	fmt.Println(board)
}

const boardWidth = 8

type Board [boardWidth][boardWidth]*Piece

func NewBoard() Board {
	board := new(Board)
	for row := 0; row < 3; row++ {
		var col int
		if row%2 == 0 {
			col = 1
		} else {
			col = 0
		}
		for ; col < boardWidth; col += 2 {
			p := new(Piece)
			p.color = Red
			board[row][col] = p
		}
	}
	for row := boardWidth - 1; row > boardWidth-4; row-- {
		var col int
		if row%2 == 0 {
			col = 1
		} else {
			col = 0
		}
		for ; col < boardWidth; col += 2 {
			p := new(Piece)
			p.color = Black
			board[row][col] = p
		}
	}
	return *board
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

type Color int

const (
	Red Color = iota
	Black
)

type Kind int

const (
	Pawn Kind = iota
	King
)

type Piece struct {
	color Color
	kind  Kind
}
