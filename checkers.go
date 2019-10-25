package main

import (
	"fmt"
	"github.com/tech6hutch/checkers/model/board"
)

func main() {
	fmt.Println("Checkers!")
	board := board.New()

	fmt.Println(board)
}
