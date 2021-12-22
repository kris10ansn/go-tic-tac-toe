package main

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/src/cli"
	"github.com/kris10ansn/go-tic-tac-toe/src/game"
)

func main() {
	var winner game.Tic = game.EMPTY_TIC
	var turn game.Tic = game.X_TIC
	var moves byte = 0

	gameBoard := game.CreateEmptyBoard()
	cli.PrintBoard(gameBoard)

	for ; winner == game.EMPTY_TIC && moves < 9; moves++ {
		fmt.Printf("%s's turn\n", game.TicToString(turn))

		x, y := cli.WaitForMove(gameBoard)

		game.SetBoardCoordinate(gameBoard, x, y, turn)
		winner = game.CheckWin(gameBoard)

		cli.PrintBoard(gameBoard)
		game.NextTurn(&turn)
	}

	cli.PrintGameEnd(gameBoard, winner, moves)
}
