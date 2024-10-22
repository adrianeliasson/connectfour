package main

import (
	"fmt"
)

func main() {
	board := Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}
	game := &Game{board, 1}
	fmt.Println(game.board)
	fmt.Println("HELLO")
	play(game)
}

func play(game *Game) {
	for !game.isOver() {
		game.makeTurn()
	}
}
