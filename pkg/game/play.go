package game

import "fmt"

type Player interface {
	AwaitMove(board *Board) (int, int)
	Present(board *Board)
	EndGame(board *Board, winner Tic, moves byte)
}

type Game struct {
	playerX Player
	playerO Player
}

func (game *Game) Play() {
	var (
		board  *Board = CreateEmptyBoard()
		winner Tic    = EMPTY_TIC
		turn   Tic    = X_TIC
		moves  byte   = 0
	)

	for ; winner == EMPTY_TIC && moves < 9; moves++ {
		game.Present(board)

		x, y := game.GetPlayer(turn).AwaitMove(board)

		SetBoardCoordinate(board, x, y, turn)

		winner = CheckWin(board)
		NextTurn(&turn)
	}

	game.EndGame(board, winner, moves)
}

func (game *Game) Present(board *Board) {
	game.playerX.Present(board)
	game.playerO.Present(board)
}

func (game *Game) EndGame(board *Board, winner Tic, moves byte) {
	game.playerX.EndGame(board, winner, moves)
	game.playerO.EndGame(board, winner, moves)
}

func (game *Game) GetPlayer(tic Tic) Player {
	if tic == X_TIC {
		return game.playerX
	} else if tic == O_TIC {
		return game.playerO
	} else {
		panic(fmt.Sprintf("Unknown tic: %d", tic))
	}
}
