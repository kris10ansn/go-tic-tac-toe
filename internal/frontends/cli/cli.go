package cli

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

type CLIFrontEnd struct {
	game.FrontEnd
}

func (CLIFrontEnd) AwaitMove(board *game.Board, turn game.Tic) (int, int) {
	fmt.Printf("%s's turn\n", game.TicToString(turn))

	var (
		x   int
		y   int
		err error
	)

	for {
		if x, y, err = inputCoordinates(board); err != nil {
			fmt.Printf("%s\n", err)
		} else {
			break
		}
	}

	return x, y
}

func (CLIFrontEnd) PresentBoard(board *game.Board) {
	fmt.Println(game.BoardToString(board))
}

func (CLIFrontEnd) EndGame(board *game.Board, winner game.Tic, moves byte) {
	if winner != game.EMPTY_TIC {
		fmt.Printf("%s won the game after %d moves!\n", game.TicToString(winner), moves)
	} else if moves == 9 {
		fmt.Println("It's a draw!")
	} else {
		panic("Unexpected behaviour")
	}
}

func inputCoordinates(board *game.Board) (int, int, error) {
	var (
		x   int
		y   int
		err error
	)

	fmt.Print("Row number (0-2):")
	_, err = fmt.Scanf("%d\n", &y)

	if err != nil {
		return 0, 0, err
	}

	fmt.Print("Column number (0-2):")
	_, err = fmt.Scanf("%d\n", &x)

	if err != nil {
		return 0, 0, err
	}

	if x > 2 || y > 2 {
		return x, y, fmt.Errorf("[%d, %d] coordinates out of bounds", x, y)
	}

	if board[y][x] != game.EMPTY_TIC {
		return x, y, fmt.Errorf("[%d, %d] is already occupied", x, y)
	}

	return x, y, nil
}
