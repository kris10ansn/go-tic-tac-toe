package cli

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

type CommandlineInterface struct {
	boardPresents     int
	numberOfPlayers   int
	endMessagePrinted bool
}

func (cli *CommandlineInterface) CreateCommandlinePlayer() *CommandlinePlayer {
	cli.numberOfPlayers++

	return &CommandlinePlayer{
		CLI: cli,
	}
}

func (cli *CommandlineInterface) present(board *game.Board) {
	// Prevent multiple board prints after each move
	if cli.boardPresents%cli.numberOfPlayers == 0 {
		fmt.Println(game.BoardToString(board))
	}

	cli.boardPresents++
}

func (cli *CommandlineInterface) printTurn(turn game.Tic) {
	if cli.numberOfPlayers == 1 {
		fmt.Println("Your turn!")
	} else {
		fmt.Printf("%s's turn!\n", game.TicToString(turn))
	}
}

func (cli *CommandlineInterface) endGame(winner game.Tic, moves byte) {
	if !cli.endMessagePrinted {
		if winner == game.EMPTY_TIC {
			fmt.Println("It's a draw!")
		} else {
			fmt.Printf(
				"%s won the game after %d moves!\n",
				game.TicToString(winner),
				moves,
			)
		}

		cli.endMessagePrinted = true
	}
}
