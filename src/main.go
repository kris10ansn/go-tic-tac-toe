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
		fmt.Printf("%c's turn\n", turn)

		x, y := cli.WaitForMove(gameBoard)

		gameBoard[y][x] = turn
		winner = game.CheckWin(gameBoard)

		cli.PrintBoard(gameBoard)
		game.NextTurn(&turn)
	}

	if winner != game.EMPTY_TIC {
		fmt.Printf("%c won the game!\n", winner)
	} else if moves == 9 {
		fmt.Println("It's a draw!")
	} else {
		panic("Unexpected behaviour")
	}
}
