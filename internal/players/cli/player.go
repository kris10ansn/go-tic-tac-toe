package cli

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

type commandlinePlayer struct {
	game.Player
	manager *CommandlineManager
}

func (p *commandlinePlayer) AwaitMove(board *game.Board, turn game.Tic) (int, int) {
	p.manager.printTurn(turn)

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

func (p *commandlinePlayer) Present(board *game.Board) {
	p.manager.present(board)
}
func (p *commandlinePlayer) EndGame(winner game.Tic, moves byte) {
	p.manager.endGame(winner, moves)
}
