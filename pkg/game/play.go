package game

import "fmt"

type Player interface {
	AwaitMove(board *Board, turn Tic) (int, int)
	Present(board *Board)
	EndGame(winner Tic, moves byte)
}

type Game struct {
	playerX Player
	playerO Player
}

func CreateGame(playerX Player, playerO Player) *Game {
	return &Game{
		playerX: playerX,
		playerO: playerO,
	}
}

func (game *Game) Play() {
	var (
		board  *Board = createEmptyBoard()
		winner Tic    = EMPTY_TIC
		turn   Tic    = X_TIC
		moves  byte   = 0
	)

	for ; winner == EMPTY_TIC && moves < 9; moves++ {
		game.Present(board)

		x, y := game.GetPlayer(turn).AwaitMove(board, turn)

		setBoardCoordinate(board, x, y, turn)

		winner = checkWin(board)
		nextTurn(&turn)
	}

	game.Present(board)
	game.EndGame(winner, moves)
}

func (game *Game) Present(board *Board) {
	game.playerX.Present(board)
	game.playerO.Present(board)
}

func (game *Game) EndGame(winner Tic, moves byte) {
	game.playerX.EndGame(winner, moves)
	game.playerO.EndGame(winner, moves)
}

func (game *Game) player(tic Tic) *Player {
	if tic == X_TIC {
		return &game.playerX
	} else if tic == O_TIC {
		return &game.playerO
	} else {
		panic(fmt.Sprintf("Unknown tic: %d", tic))
	}
}

func (game *Game) GetPlayer(tic Tic) Player {
	return *game.player(tic)
}

func (game *Game) SetPlayer(tic Tic, player Player) {
	*game.player(tic) = player
}

func (game *Game) AddPlayer(player Player) (Tic, error) {
	if *game.player(X_TIC) == nil {
		game.SetPlayer(X_TIC, player)
		return X_TIC, nil
	} else if *game.player(O_TIC) == nil {
		game.SetPlayer(O_TIC, player)
		return O_TIC, nil
	}

	return EMPTY_TIC, fmt.Errorf("Game full")
}
