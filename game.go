package main

import (
	"math/rand"
	"strings"
)

func (board Board) rollDice() Board {
	board.uses = 0
	board.dice1 = rand.Intn(6) + 1
	board.dice2 = rand.Intn(6) + 1
	board.adice1 = board.dice1
	board.adice2 = board.dice2

	// board = board.checkPossibility()

	if board.dice1 == board.dice2 {
		board.uses = 4
	}
	return board
}

func (board Board) move(character string, diceint int) (Board, string) {
	var funcerr string
	character = strings.ToUpper(character)
	if diceint != board.dice1 && diceint != board.dice2 {
		funcerr = "You don't have such number!"
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
		funcerr = "That character doesn't exist!"
		return board, funcerr
	}

	if board.table[charint].color != board.turn {
		funcerr = "There are no rocks of your color on that slot!"
		return board, funcerr
	}

	if board.turn == 0 {
		if charint+diceint > 23 {
			funcerr = "A move there isn't possible!"
			return board, funcerr
		} else if board.table[charint+diceint].color == board.turn || board.table[charint+diceint].color == -1 || board.table[charint+diceint].amount == 1 {
			board.table[charint].amount--
			if board.table[charint].amount == 0 {
				board.table[charint].color = -1
			}

			if board.table[charint+diceint].color == 1 {
				board.holding[1]++
				board.table[charint+diceint].amount = 0
			}

			board.table[charint+diceint].amount++
			board.table[charint+diceint].color = 0

			if board.uses == 0 {
				if diceint == board.dice1 {
					board.adice1 = -1
				} else {
					board.adice2 = -1
				}
			} else {
				board.uses--
			}
		} else {
			funcerr = "A move there isn't possible!"
			return board, funcerr
		}
	} else {
		if charint-diceint < 0 {
			funcerr = "A move there isn't possible!"
			return board, funcerr
		} else if board.table[charint-diceint].color == board.turn || board.table[charint-diceint].color == -1 || board.table[charint-diceint].amount == 1 {
			board.table[charint].amount--
			if board.table[charint].amount == 0 {
				board.table[charint].color = -1
			}

			if board.table[charint-diceint].color == 0 {
				board.holding[0]++
				board.table[charint-diceint].amount = 0
			}

			board.table[charint-diceint].amount++
			board.table[charint-diceint].color = 1

			if board.uses == 0 {
				if diceint == board.dice1 {
					board.adice1 = -1
				} else {
					board.adice2 = -1
				}
			} else {
				board.uses--
			}
		} else {
			funcerr = "A move there isn't possible!"
			return board, funcerr
		}
	}

	return board, funcerr
}

func (board Board) checkIfSittingPossible() bool {
	if board.turn == 0 {
		if (board.table[board.dice1-1].color == board.turn || board.table[board.dice1-1].color == -1 || board.table[board.dice1-1].amount == 1) && board.adice1 != -1 {
			return true
		} else if (board.table[board.dice2-1].color == board.turn || board.table[board.dice2-1].color == -1 || board.table[board.dice2-1].amount == 1) && board.adice2 != -1 {
			return true
		} else {
			return false
		}
	} else {
		if (board.table[24-board.dice1].color == board.turn || board.table[24-board.dice1].color == -1 || board.table[24-board.dice1].amount == 1) && board.adice1 != -1 {
			return true
		} else if (board.table[24-board.dice2].color == board.turn || board.table[24-board.dice2].color == -1 || board.table[24-board.dice2].amount == 1) && board.adice2 != -1 {
			return true
		} else {
			return false
		}
	}
}

func (board Board) place(diceint int) (Board, string) {
	var funcerr string
	if diceint != board.dice1 && diceint != board.dice2 {
		funcerr = "You don't have such number!"
		return board, funcerr
	}
	if board.turn == 0 {
		if board.table[diceint-1].color == board.turn || board.table[diceint-1].color == -1 || board.table[diceint-1].amount == 1 {
			board.holding[0]--
			if board.table[diceint-1].color == 1 {
				board.holding[1]++
				board.table[diceint-1].amount = 0
			}

			board.table[diceint-1].amount++
			board.table[diceint-1].color = 0

			if board.uses == 0 {
				if diceint == board.dice1 {
					board.adice1 = -1
				} else {
					board.adice2 = -1
				}
			} else {
				board.uses--
			}
			return board, funcerr
		}
		funcerr = "You cannot place your rock there!"
		return board, funcerr
	}
	if board.table[24-diceint].color == board.turn || board.table[24-diceint].color == -1 || board.table[24-diceint].amount == 1 {
		board.holding[1]--
		if board.table[24-diceint].color == 0 {
			board.holding[0]++
			board.table[24-diceint].amount = 0
		}

		board.table[24-diceint].amount++
		board.table[24-diceint].color = 1

		if board.uses == 0 {
			if diceint == board.dice1 {
				board.adice1 = -1
			} else {
				board.adice2 = -1
			}
		} else {
			board.uses--
		}
		return board, funcerr
	}
	funcerr = "You cannot place your rock there!"
	return board, funcerr
}
