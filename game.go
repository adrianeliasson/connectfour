package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func transpose(a [][]int) [][]int {
	newArr := make([][]int, len(a))
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}

	return newArr
}

type Board [][]int
type Game struct {
	board      Board
	playerTurn int
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

func count[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
}
func countOnesOrTwos(slice []int) int {
	return count(slice, func(x int) bool { return x == 1 || x == 2 })
}
func (g *Game) placePiece(colNumber int) {
	numberOfNonZeroes := countOnesOrTwos(g.board[colNumber])
	if len(g.board[colNumber]) < numberOfNonZeroes+1 {
		fmt.Println("That column is full")
		return
	}
	g.board[colNumber][numberOfNonZeroes] = g.playerTurn
	g.changeTurn()
}

func (g Game) isOver() bool {
	return false
}

func (g *Game) changeTurn() {
	fmt.Println("Current: ", g.playerTurn)
	if g.playerTurn == 1 {
		g.playerTurn = 2
	} else {
		g.playerTurn = 1
	}
}

func (g Game) printGameState() {
	for i := 6; i >= 0; i-- {
		for j := 0; j < 7; j++ {
			fmt.Print(xoro(g.board[j][i]), " ")
		}
		fmt.Print("\n")
	}
}

func (g *Game) makeTurn() {
	g.printGameState()
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
