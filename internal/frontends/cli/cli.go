package cli

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

type CLIFrontEnd struct{}

func (CLIFrontEnd) AwaitMove(board *game.Board, turn game.Tic) (byte, byte) {
	fmt.Printf("%s's turn\n", game.TicToString(turn))

	var (
		x   byte
		y   byte
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

func New() game.FrontEnd {
	return CLIFrontEnd{}
}

func inputCoordinates(board *game.Board) (byte, byte, error) {
	var (
		x   byte
		y   byte
		err error
	)

	y, err = promptByteInput("Row number (0-2):")

	if err != nil {
		return 0, 0, err
	}

	x, err = promptByteInput("Column number (0-2): ")

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

func inputByte() (byte, error) {
	var result byte
	_, err := fmt.Scanf("%d\n", &result)

	return result, err
}

func promptByteInput(message string) (byte, error) {
	fmt.Print(message)
	return inputByte()
}
