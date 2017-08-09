package main

import (
	"math/rand"
	"time"
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
	adice1  int
	dice2   int
	adice2  int
	game    bool
	uses    int
}

func (board Board) initiateBoard() Board {
	rand.Seed(time.Now().UnixNano())
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
