package main

import (
	"fmt"
	"strconv"
)

const escape = "\x1b"
const white = escape + "[37m"
const black = escape + "[30m"
const cyanB = escape + "[46m"
const blue = escape + "[34m"
const yellowB = escape + "[43m"
const reset = escape + "[0m"

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
	ro := yellowB + blue + " A B C D E F  G H I J K L \n"
	// Top half
	for i := 1; i < 6; i++ {
		ro += cyanB + " "
		for y := 0; y < 6; y++ {
			if board.table[y].amount >= i {
				if board.table[y].color == 1 {
					ro += white + "●"
				} else {
					ro += black + "●"
				}
			} else {
				ro += " "
			}
			ro += " "
		}
		ro += " "
		for y := 6; y < 12; y++ {
			if board.table[y].amount >= i {
				if board.table[y].color == 1 {
					ro += white + "●"
				} else {
					ro += black + "●"
				}
			} else {
				ro += " "
			}
			ro += " "
		}
		ro += "" + reset + "\n"
	}
	// Middle
	for i := 0; i < 6; i++ {
		ro += cyanB + white + " "
		if board.table[i].amount > 5 {
			ro += strconv.Itoa(board.table[i].amount)
		} else {
			ro += " "
		}
	}
	ro += " "
	for i := 6; i < 12; i++ {
		ro += " "
		if board.table[i].amount > 5 {
			ro += strconv.Itoa(board.table[i].amount)
		} else {
			ro += " "
		}
	}
	ro += " " + reset + "\n" + yellowB + "                          " + reset + "\n" + cyanB + white
	for i := 23; i > 17; i-- {
		ro += " "
		if board.table[i].amount > 5 {
			ro += strconv.Itoa(board.table[i].amount)
		} else {
			ro += " "
		}
	}
	ro += " "
	for i := 17; i > 11; i-- {
		ro += " "
		if board.table[i].amount > 5 {
			ro += strconv.Itoa(board.table[i].amount)
		} else {
			ro += " "
		}
	}
	ro += " " + reset + "\n" + cyanB
	// Bottom Half
	for i := 5; i > 0; i-- {
		ro += " "
		for y := 23; y > 17; y-- {
			if board.table[y].amount >= i {
				if board.table[y].color == 1 {
					ro += white + "●"
				} else {
					ro += black + "●"
				}
			} else {
				ro += " "
			}
			ro += " "
		}
		ro += " "
		for y := 17; y > 11; y-- {
			if board.table[y].amount >= i {
				if board.table[y].color == 1 {
					ro += white + "●"
				} else {
					ro += black + "●"
				}
			} else {
				ro += " "
			}
			ro += " "
		}
		ro += reset + "\n" + cyanB
	}
	ro += yellowB + blue + " M N O P Q R  S T U V W X " + reset + "\n"
	fmt.Println(ro)
}
