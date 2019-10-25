package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Checkers!")
	board := new(Board)
	for row := 0; row < 3; row++ {
		var col int
		if row%2 == 0 {
			col = 0
		} else {
			col = 1
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
			col = 0
		} else {
			col = 1
		}
		for ; col < boardWidth; col += 2 {
			p := new(Piece)
			p.color = Black
			board[row][col] = p
		}
	}

	
}

const boardWidth = 8

type Board [boardWidth][boardWidth]*Piece

func String(board Board) string {
	lines := new([boardWidth][boardWidth + 1]rune)
	for row := range board {
		for col := range board[row] {
			p := board[row][col]
			if p == nil {
				lines[row][col] = '#'
			} else {
				lines[row][col] = '*'
			}
		}
		lines[row][boardWidth] = '\n'
	}
	s := new(string)
	for line := range lines {
		*s += string(line)
	}
	return *s
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
