package main

import (
	"bufio"
	"fmt"
	"github.com/tech6hutch/checkers/model/board"
	"github.com/tech6hutch/checkers/model/pieces"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Checkers!")
	board := board.New()
	playerColors := &map[int]pieces.Color{
		1: pieces.ColorDark,
		2: pieces.ColorLight,
	}
	playerTurn := 1

	for {
		switch playerTurn {
		case 1: playerTurn = 2
		case 2: playerTurn = 1
		default: panic("wtf")
		}

		fmt.Printf("Player %d's turn! (%s)\n",
			playerTurn, (*playerColors)[playerTurn])
		fmt.Println(board)

		scanner.Scan()
		line := scanner.Text()
		fmt.Println(line)
	}

	fmt.Println(board)
}
