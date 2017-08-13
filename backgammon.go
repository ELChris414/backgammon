package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	color.New(color.BgCyan)
	fmt.Println("Welcome to Backgammon!")
	board := Board{}
	board = board.initiateBoard()
	var err string
	for board.game == true {
		board = board.rollDice()
		board.render()
		if board.holding[board.turn] == 0 {
			board = board.attemptMove()
		} else {
			if !board.checkIfSittingPossible() {
				fmt.Println("Sadly you cannot play!")
				if board.turn == 0 {
					board.turn = 1
				} else {
					board.turn = 0
				}
				continue
			}
			for board.adice1 != -1 || board.adice2 != -1 || board.uses > 0 {
				if board.holding[board.turn] > 0 {
					reader := bufio.NewReader(os.Stdin)
					fmt.Println("place *num*")
					text, _ := reader.ReadString('\n')
					stuff := strings.Fields(text)
					if len(stuff) == 2 {
						diceint, interr := strconv.Atoi(stuff[1])
						if interr != nil {
							fmt.Println("\nYou didn't give a number!")
							board.render()
							continue
						}
						if stuff[0] == "place" {
							board, err = board.place(diceint)
							if err != "" {
								fmt.Println(err)
								board.render()
								continue
							}
						}
					} else {
						fmt.Println("You didn't give enough arguments!")
						board.render()
						continue
					}
					fmt.Println()
					if board.adice1 != -1 || board.adice2 != -1 {
						board.render()
					}
				} else {
					board = board.attemptMove()
				}
			}
		}
	}
}

func (board Board) attemptMove() Board {
	var err string
	if board.uses == 0 {
		for board.adice1 != -1 || board.adice2 != -1 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("move *character* *num*")
			text, _ := reader.ReadString('\n')
			stuff := strings.Fields(text)
			if len(stuff) == 3 {
				diceint, interr := strconv.Atoi(stuff[2])
				if interr != nil {
					fmt.Println("You didn't give a number!")
					board.render()
					continue
				}
				if stuff[0] == "move" {
					board, err = board.move(stuff[1], diceint)
					if err != "" {
						fmt.Println(err)
					}
				}
			} else {
				fmt.Println("You didn't give enough arguments!")
				board.render()
				continue
			}
			fmt.Println()
			if board.adice1 != -1 || board.adice2 != -1 {
				board.render()
			}
		}
	} else {
		for board.uses != 0 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("move *character*")
			text, _ := reader.ReadString('\n')
			stuff := strings.Fields(text)
			if len(stuff) == 2 && stuff[0] == "move" {
				board, err = board.move(stuff[1], board.dice1)
				if err != "" {
					fmt.Println(err)
				}
			}
			fmt.Println()
			if board.uses != 0 {
				board.render()
			}
		}
	}
	if board.turn == 0 {
		board.turn = 1
	} else {
		board.turn = 0
	}
	return board
}
