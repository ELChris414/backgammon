package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// Slot struct
type Slot struct {
	amount, color int
}

// Board struct
type Board struct {
	table   map[int]*Slot
	turn    int
	holding map[int]int
	dice1   int
	dice2   int
	game    bool
	uses    int
}

func (board Board) initiateBoard() Board {
	board.turn = 0
	board.game = true
	board.table = make(map[int]*Slot)
	board.holding = make(map[int]int)
	board.holding[0] = -1
	board.holding[1] = -1
	board.table[0] = &Slot{2, 0}
	board.table[1] = &Slot{0, -1}
	board.table[2] = &Slot{0, -1}
	board.table[3] = &Slot{0, -1}
	board.table[4] = &Slot{0, -1}
	board.table[5] = &Slot{5, 1}
	board.table[6] = &Slot{0, -1}
	board.table[7] = &Slot{3, 1}
	board.table[8] = &Slot{0, -1}
	board.table[9] = &Slot{0, -1}
	board.table[10] = &Slot{0, -1}
	board.table[11] = &Slot{5, 0}
	board.table[12] = &Slot{5, 1}
	board.table[13] = &Slot{0, -1}
	board.table[14] = &Slot{0, -1}
	board.table[15] = &Slot{0, -1}
	board.table[16] = &Slot{3, 0}
	board.table[17] = &Slot{0, -1}
	board.table[18] = &Slot{5, 0}
	board.table[19] = &Slot{0, -1}
	board.table[20] = &Slot{0, -1}
	board.table[21] = &Slot{0, -1}
	board.table[22] = &Slot{0, -1}
	board.table[23] = &Slot{2, 1}
	return board
}

func (board Board) render() {
	fmt.Print("Current turn: ")
	if board.turn == 1 {
		fmt.Print("White")
	} else {
		fmt.Print("Black")
	}
	if board.dice1 != -1 && board.dice2 != -1 && board.uses == 0 {
		fmt.Println("\nThe dice rolled: " + strconv.Itoa(board.dice1) + " " + strconv.Itoa(board.dice2))
	} else if board.uses != 0 {
		fmt.Println("\nThe dice rolled: " + strconv.Itoa(board.dice1) + " " + strconv.Itoa(board.dice2) + "! You have " + strconv.Itoa(board.uses) + " more uses!")
	} else if board.dice1 == -1 {
		fmt.Println("\nYou still have to play : " + strconv.Itoa(board.dice2))
	} else if board.dice2 == -1 {
		fmt.Println("\nYou still have to play : " + strconv.Itoa(board.dice1))
	}
	white := color.New(color.BgCyan, color.FgWhite).PrintFunc()
	black := color.New(color.BgCyan, color.FgBlack).PrintFunc()
	blue := color.New(color.FgBlue, color.BgYellow).PrintFunc()
	blue(" A B C D E F  G H I J K L ")
	fmt.Println()
	// Top half
	for i := 1; i < 6; i++ {
		white(" ")
		for y := 0; y < 6; y++ {
			if board.table[y].amount >= i {
				if board.table[y].color == 1 {
					white("●")
				} else {
					black("●")
				}
			} else {
				white(" ")
			}
			white(" ")
		}
		white(" ")
		for y := 6; y < 12; y++ {
			if board.table[y].amount >= i {
				if board.table[y].color == 1 {
					white("●")
				} else {
					black("●")
				}
			} else {
				white(" ")
			}
			white(" ")
		}
		white("")
		fmt.Println()
	}
	// Middle
	for i := 0; i < 6; i++ {
		white(" ")
		if board.table[i].amount > 5 {
			white(board.table[i].amount)
		} else {
			white(" ")
		}
	}
	white(" ")
	for i := 6; i < 12; i++ {
		white(" ")
		if board.table[i].amount > 5 {
			white(board.table[i].amount)
		} else {
			white(" ")
		}
	}
	white(" ")
	fmt.Println()
	blue("                          ")
	fmt.Println()
	for i := 23; i > 17; i-- {
		white(" ")
		if board.table[i].amount > 5 {
			white(board.table[i].amount)
		} else {
			white(" ")
		}
	}
	white(" ")
	for i := 17; i > 11; i-- {
		white(" ")
		if board.table[i].amount > 5 {
			white(board.table[i].amount)
		} else {
			white(" ")
		}
	}
	white(" ")
	fmt.Println()
	// Bottom Half
	for i := 5; i > 0; i-- {
		white(" ")
		for y := 23; y > 17; y-- {
			if board.table[y].amount >= i {
				if board.table[y].color == 1 {
					white("●")
				} else {
					black("●")
				}
			} else {
				white(" ")
			}
			white(" ")
		}
		white(" ")
		for y := 17; y > 11; y-- {
			if board.table[y].amount >= i {
				if board.table[y].color == 1 {
					white("●")
				} else {
					black("●")
				}
			} else {
				white(" ")
			}
			white(" ")
		}
		fmt.Println()
	}
	blue(" M N O P Q R  S T U V W X ")
	fmt.Println()
}

func (board Board) rollDice() Board {
	board.uses = 0
	board.dice1 = rand.Intn(5) + 1
	board.dice2 = rand.Intn(5) + 1
	if board.dice1 == board.dice2 {
		board.uses = 4
	}
	return board
}

func (board Board) move(character string, dice string) (Board, error) {
	character = strings.ToUpper(character)
	diceint, err := strconv.Atoi(dice)
	funcerr := errors.New("")
	if err != nil {
		funcerr = errors.New("You didn't give a number!")
		return board, funcerr
	}
	if diceint != board.dice1 && diceint != board.dice2 {
		funcerr = errors.New("You don't have such number!")
	}
	charint := -1
	switch character {
	case "A":
		charint = 0
		break
	case "B":
		charint = 1
		break
	case "C":
		charint = 2
		break
	case "D":
		charint = 3
		break
	case "E":
		charint = 4
		break
	case "F":
		charint = 5
		break
	case "G":
		charint = 6
		break
	case "H":
		charint = 7
		break
	case "I":
		charint = 8
		break
	case "J":
		charint = 9
		break
	case "K":
		charint = 10
		break
	case "L":
		charint = 11
		break
	case "X":
		charint = 12
		break
	case "W":
		charint = 13
		break
	case "V":
		charint = 14
		break
	case "U":
		charint = 15
		break
	case "T":
		charint = 16
		break
	case "S":
		charint = 17
		break
	case "R":
		charint = 18
		break
	case "Q":
		charint = 19
		break
	case "P":
		charint = 20
		break
	case "O":
		charint = 21
		break
	case "N":
		charint = 22
		break
	case "M":
		charint = 23
		break
	default:
		funcerr = errors.New("That character doesn't exist!")
		return board, funcerr
	}

	if board.table[charint].color != board.turn {
		funcerr = errors.New("There are no rocks of your color on that slot!")
		return board, funcerr
	}

	if board.table[charint+diceint].color == board.turn || board.table[charint+diceint].color == -1 || board.table[charint+diceint].amount == 1 {
		board.table[charint].amount--
		if board.table[charint].amount == 0 {
			board.table[charint].color = -1
		}

		if board.table[charint+diceint].color == 1 {
			board.holding[1]++
		}

		board.table[charint+diceint].amount++
		board.table[charint+diceint].color = 0

		if board.uses == 0 {
			if diceint == board.dice1 {
				board.dice1 = -1
			} else {
				board.dice2 = -1
			}
		} else {
			board.uses--
		}
	} else {
		funcerr = errors.New("A move there isn't possible!")
		return board, funcerr
	}

	return board, funcerr
}
