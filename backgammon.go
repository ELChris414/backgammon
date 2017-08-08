package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Backgammon!")
	board := Board{}
	board.initiateBoard()
	board.render()
}
