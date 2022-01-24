package random

import (
	"math/rand"

	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

type RandomPlayer struct{}

func (RandomPlayer) AwaitMove(board *game.Board, turn game.Tic) (int, int) {
	for {
		x := rand.Intn(3)
		y := rand.Intn(3)

		if board[y][x] == game.EMPTY_TIC {
			return x, y
		}
	}
}

func (RandomPlayer) Present(board *game.Board)           {}
func (RandomPlayer) EndGame(winner game.Tic, moves byte) {}
