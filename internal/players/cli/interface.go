package cli

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

type CommandlineManager struct {
	boardPresents     int
	numberOfPlayers   int
	endMessagePrinted bool
}

func (cli *CommandlineManager) CreateCommandlinePlayer() *commandlinePlayer {
	cli.numberOfPlayers++

	return &commandlinePlayer{
		manager: cli,
	}
}

func (m *CommandlineManager) present(board *game.Board) {
	// Prevent multiple board prints after each move
	if m.boardPresents%m.numberOfPlayers == 0 {
		fmt.Println(game.BoardToString(board))
	}

	m.boardPresents++
}

func (m *CommandlineManager) printTurn(turn game.Tic) {
	if m.numberOfPlayers == 1 {
		fmt.Println("Your turn!")
	} else {
		fmt.Printf("%s's turn!\n", game.TicToString(turn))
	}
}

func (m *CommandlineManager) endGame(winner game.Tic, moves byte) {
	if !m.endMessagePrinted {
		if winner == game.EMPTY_TIC {
			fmt.Println("It's a draw!")
		} else {
			fmt.Printf(
				"%s won the game after %d moves!\n",
				game.TicToString(winner),
				moves,
			)
		}

		m.endMessagePrinted = true
	}
}
