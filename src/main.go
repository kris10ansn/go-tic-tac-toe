package main

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/src/cli"
	"github.com/kris10ansn/go-tic-tac-toe/src/game"
)

func main() {
	var turn game.Tic = game.X_TIC
	var moves byte = 0

	gameBoard := game.CreateEmptyBoard()
	cli.PrintBoard(gameBoard)

	for ; moves < 9; moves++ {
		fmt.Printf("%c's turn\n", turn)

		x, y := cli.WaitForMove(gameBoard)

		gameBoard[y][x] = turn

		// TODO: check win

		cli.PrintBoard(gameBoard)

		game.NextTurn(&turn)
	}

	// TODO: draw
}
