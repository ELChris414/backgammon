package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	color.New(color.BgCyan)
	fmt.Println("Welcome to Backgammon!")
	board := Board{}
	board = board.initiateBoard()
	var err error
	/*for board.game == true {
		board = board.rollDice()
		board.render()
		if board.uses == -1 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("move *character* *num*")
			text, _ := reader.ReadString('\n')
			stuff := strings.Fields(text)
			if len(stuff) == 3 && stuff[0] == "move" {
				board, err = board.move(stuff[1], stuff[2])
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {

		}
	}*/
	board = board.rollDice()
	board.render()
	if board.uses == 0 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("move *character* *num*")
		text, _ := reader.ReadString('\n')
		stuff := strings.Fields(text)
		if len(stuff) == 3 && stuff[0] == "move" {
			board, err = board.move(stuff[1], stuff[2])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	board.render()
}
