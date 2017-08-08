package main

import (
	"fmt"

	"github.com/fatih/color"
)

// Slot struct
type Slot struct {
	amount, color int
}

// Board struct
type Board struct {
	table   map[int]Slot
	turn    bool
	holding map[int]Slot
}

func (board Board) initiateBoard() {
	board.table[0] = Slot{2, 0}
	board.table[1] = Slot{0, -1}
	board.table[2] = Slot{0, -1}
	board.table[3] = Slot{0, -1}
	board.table[4] = Slot{0, -1}
	board.table[5] = Slot{5, 1}
	board.table[6] = Slot{0, -1}
	board.table[7] = Slot{3, 1}
	board.table[8] = Slot{0, -1}
	board.table[9] = Slot{0, -1}
	board.table[10] = Slot{0, -1}
	board.table[11] = Slot{5, 0}
	board.table[12] = Slot{5, 1}
	board.table[13] = Slot{0, -1}
	board.table[14] = Slot{0, -1}
	board.table[15] = Slot{0, -1}
	board.table[16] = Slot{3, 0}
	board.table[17] = Slot{0, -1}
	board.table[18] = Slot{5, 0}
	board.table[19] = Slot{0, -1}
	board.table[20] = Slot{0, -1}
	board.table[21] = Slot{0, -1}
	board.table[22] = Slot{0, -1}
	board.table[23] = Slot{2, 1}
}

func (board Board) render() {
	fmt.Print("Current turn: ")
	if board.turn == true {
		fmt.Print("White")
	} else {
		fmt.Print("Black")
	}
	fmt.Println()
	color.Blue(" A B C D E F   G H I J K L")
	for i := 0; i < 5; i++ {
		fmt.Print(" ")
		for y := 0; y < 6; y++ {
			if board.table[y].amount >= i {
				if board.table[y].color == 1 {
					color.White("●")
				} else {
					color.Black("●")
				}
			}
			fmt.Print(" ")
		}
		fmt.Print("  ")
		for y := 6; y < 12; y++ {
			if board.table[y].amount >= i {
				if board.table[y].color == 1 {
					color.White("●")
				} else {
					color.Black("●")
				}
			}
			fmt.Print(" ")
		}
	}

	color.Blue(" M N O P Q R   S T U V W X")

}
