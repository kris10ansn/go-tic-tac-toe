package cli

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

type CommandlinePlayer struct {
	game.Player
	CLI *CommandlineInterface
}

func (p *CommandlinePlayer) AwaitMove(board *game.Board, turn game.Tic) (int, int) {
	p.CLI.printTurn(turn)

	var x int
	var y int
	var err error

	for {
		if x, y, err = inputCoordinates(board); err != nil {
			fmt.Printf("%s\n", err)
			continue
		}

		break
	}

	return x, y
}

func (p *CommandlinePlayer) Present(board *game.Board) {
	p.CLI.present(board)
}
func (p *CommandlinePlayer) EndGame(winner game.Tic, moves byte) {
	p.CLI.endGame(winner, moves)
}
