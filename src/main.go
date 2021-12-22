package main

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/src/cli"
	"github.com/kris10ansn/go-tic-tac-toe/src/game"
)

func main() {
	var (
		gameBoard game.Board = game.CreateEmptyBoard()
		winner    game.Tic   = game.EMPTY_TIC
		turn      game.Tic   = game.X_TIC
		moves     byte       = 0
	)

	cli.PrintBoard(gameBoard)

	for ; winner == game.EMPTY_TIC && moves < 9; moves++ {
		fmt.Printf("%s's turn\n", game.TicToString(turn))

		x, y := cli.WaitForMove(gameBoard)

		game.SetBoardCoordinate(&gameBoard, x, y, turn)
		cli.PrintBoard(gameBoard)

		winner = game.CheckWin(gameBoard)
		game.NextTurn(&turn)
	}

	cli.PrintGameEnd(gameBoard, winner, moves)
}
