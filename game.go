package main

import (
	"slices"
)

const NumberOfColumns = 7
const ColumnSize = 6

type Column []int
type Board []Column
type Game struct {
	board      []Column
	playerTurn int
}

func (g *Game) init() {
	board := make([]Column, NumberOfColumns)
	for i := range board {
		board[i] = make(Column, ColumnSize)
	}
	g.board = board
	g.playerTurn = 1
}

func (g *Game) hasEmptySlot(colNumber int) (bool, int) {
	firstEmptySlot := slices.IndexFunc(g.board[colNumber], func(n int) bool { return n == 0 })
	hasEmptySlot := firstEmptySlot >= 0
	return hasEmptySlot, firstEmptySlot
}

func (g *Game) placePiece(colNumber int) bool {
	hasEmptySlot, firstEmptySlot := g.hasEmptySlot(colNumber)
	if hasEmptySlot == false {
		return false // Fail
	}
	g.board[colNumber][firstEmptySlot] = g.playerTurn
	g.changeTurn()
	return true // Success
}

func (g Game) isOver() (bool, int) {
	// Game ends if:
	// - A player has pieces arranged in slots as
	//   - Vertical 4
	winVertical := func(game Game, player int) bool {
		consecutives := 0
		for _, column := range game.board {
			for _, slot := range column {
				if slot == player {
					consecutives++
				} else {
					consecutives = 0
				}

				if consecutives == 4 {
					return true
				}
			}
			consecutives = 0
		}
		return false
	}

	//   - Horizontal 4
	winHorizontal := func(game Game, player int) bool {
		consecutives := 0
		for i := 0; i < len(game.board[0])-1; i++ {
			for _, column := range game.board {
				if column[i] == player {
					consecutives++
				} else {
					consecutives = 0
				}
				if consecutives == 4 {
					return true
				}
			}
			consecutives = 0
		}
		return false
	}
	checkDiagonalPositive := func(game Game, player int) bool {
		walkDiagLeftToDown := func(game Game, player int, rowIndex int) bool {
			consecutives := 0
			for j := 0; j < len(game.board[0])-1; j++ {
				if rowIndex+j > len(game.board[0])-1 {
					return false
				}
				if game.board[rowIndex+j][j] == player {
					consecutives++
				} else {
					consecutives = 0
				}
				if consecutives == 4 {
					return true
				}
			}
			consecutives = 0
			return false
		}
		walkDiagTopToDown := func(game Game, player int, colIndex int) bool {
			consecutives := 0
			for i := 0; i < len(game.board[0])-1; i++ {
				if i+colIndex > len(game.board[0])-1 {
					return false
				}
				if game.board[i][colIndex+i] == player {
					consecutives++
				} else {
					consecutives = 0
				}
				if consecutives == 4 {
					return true
				}
			}
			return false
		}

		// Walk from j = 0 until board edge
		for i := 0; i < len(game.board[0])-1; i++ {
			// Walk the diag from left edge down
			if walkDiagLeftToDown(game, player, i) {
				return true
			}
		}
		for j := 0; j < len(game.board)-1; j++ {
			// Walk the diag from top edge down
			if walkDiagTopToDown(game, player, j) {
				return true
			}
		}
		return false
	}
	checkDiagonalNegative := func(game Game, player int) bool {
		walkDiagLeftToTop := func(game Game, player int, rowIndex int) bool {
			consecutives := 0
			for j := 0; j < len(game.board)-1; j++ {
				if rowIndex-j < 0 { // top edge is reached
					return false
				}
				if game.board[j][rowIndex-j] == player {
					consecutives++
				} else {
					consecutives = 0
				}
				if consecutives == 4 {
					return true
				}
			}
			consecutives = 0
			return false
		}
		walkDiagbottomToRight := func(game Game, player int, colIndex int) bool {
			consecutives := 0
			rowMaxIndex := len(game.board[0]) - 1 // 5

			for i := rowMaxIndex; i >= 0; i-- {
				currentColIndex := colIndex + rowMaxIndex - i
				if currentColIndex > rowMaxIndex+1 {
					return false
				}
				if game.board[currentColIndex][i] == player {
					consecutives++
				} else {
					consecutives = 0
				}
				if consecutives == 4 {
					return true
				}
			}
			return false
		}

		// Walk from i = 0 until board edge
		for i := 0; i < len(game.board[0])-1; i++ {
			if walkDiagLeftToTop(game, player, i) {
				return true
			}
		}
		for j := 0; j < len(game.board)-1; j++ {
			// Walk the diag from top edge down
			if walkDiagbottomToRight(game, player, j) {
				return true
			}
		}
		return false
	}

	// - Diagonal 4
	winDiagonal := func(game Game, player int) bool {
		if checkDiagonalPositive(game, player) {
			return true
		}
		if checkDiagonalNegative(game, player) {
			return true
		}
		return false
	}

	for player := 1; player <= 2; player++ {
		win := winVertical(g, player)
		if win {
			return win, player
		}
		win = winHorizontal(g, player)
		if win {
			return win, player
		}
		win = winDiagonal(g, player)
		if win {
			return win, player
		}
	}

	// - All slots are filled
	countOfZeroes := 0
	for _, c := range g.board {
		for _, m := range c {
			if m == 0 {
				countOfZeroes++
			}
		}
	}
	if countOfZeroes == 0 {
		return true, 0
	}
	return false, 0
}

func (g *Game) changeTurn() {
	if g.playerTurn == 1 {
		g.playerTurn = 2
	} else {
		g.playerTurn = 1
	}
}

func xoro(n int) string {
	if n == 0 {
		return "_"
	} else if n == 1 {
		return "X"
	} else if n == 2 {
		return "O"
	} else {
		return "WRONG"
	}
}

func (g Game) gameStateToString() string {
	stateToString := ""
	for i := range g.board[0] {
		for _, column := range g.board {
			stateToString += xoro(column[len(column)-i-1]) + " "
		}
		stateToString += "\n"
	}
	return stateToString
}
