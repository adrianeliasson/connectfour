package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
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

func (g *Game) play() {
	for !g.isOver() {
		g.gameStateToString()
		g.makeTurn()
	}
}

func (g Game) isOver() bool {
	return false
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

func (g *Game) makeTurn() {
	fmt.Println("Player ", g.playerTurn, "'s turn")
	fmt.Print("Which column do you put your piece? (or type \"quit\" to quit) [1,6]: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if strings.Compare("quit", text) == 0 {
		panic("Quitting game")
	}

	if strings.Compare("1", text) == 0 {
		g.placePiece(0)
	} else if strings.Compare("2", text) == 0 {
		g.placePiece(1)
	} else if strings.Compare("3", text) == 0 {
		g.placePiece(2)
	} else if strings.Compare("4", text) == 0 {
		g.placePiece(3)
	} else if strings.Compare("5", text) == 0 {
		g.placePiece(4)
	} else if strings.Compare("6", text) == 0 {
		g.placePiece(5)
	} else if strings.Compare("7", text) == 0 {
		g.placePiece(6)
	} else {
		fmt.Println("You didnt put a piece anywhere")
	}
}
