package main

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

func (board Board) render() {
	fmt.Print("Current turn: ")
	if board.turn == 1 {
		fmt.Print("White")
	} else {
		fmt.Print("Black")
	}
	if board.adice1 != -1 && board.adice2 != -1 && board.uses == 0 {
		fmt.Println("\nThe dice rolled: " + strconv.Itoa(board.dice1) + " " + strconv.Itoa(board.dice2))
	} else if board.uses != 0 {
		fmt.Println("\nThe dice rolled: " + strconv.Itoa(board.dice1) + " " + strconv.Itoa(board.dice2) + "! You have " + strconv.Itoa(board.uses) + " more uses!")
	} else if board.adice1 == -1 {
		fmt.Println("\nThe dice rolled: " + strconv.Itoa(board.dice1) + " " + strconv.Itoa(board.dice2) + "! You still have to play : " + strconv.Itoa(board.dice2))
	} else if board.adice2 == -1 {
		fmt.Println("\nThe dice rolled: " + strconv.Itoa(board.dice1) + " " + strconv.Itoa(board.dice2) + "! You still have to play : " + strconv.Itoa(board.dice1))
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
