package cli

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/src/game"
)

type CLIFrontEnd struct{}

func (CLIFrontEnd) AwaitMove(board game.Board, turn game.Tic) (byte, byte) {
	fmt.Printf("%s's turn\n", game.TicToString(turn))

	x, y, err := InputCoordinates(board)

	for err != nil {
		fmt.Printf("%s\n", err)
		x, y, err = InputCoordinates(board)
	}

	return x, y
}

func (CLIFrontEnd) PresentBoard(board game.Board) {
	PrintBoard(board)
}

func (CLIFrontEnd) EndGame(board game.Board, winner game.Tic, moves byte) {
	PrintBoard(board)

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

func InputByte(message string, input *byte) error {
	fmt.Print(message)
	_, err := fmt.Scanf("%d\n", input)

	return err
}

func PrintBoard(board game.Board) {
	fmt.Println(game.BoardToString(board))
}

func PrintTurn(turn game.Tic) {
	fmt.Printf("%s's turn\n", game.TicToString(turn))
}

func InputCoordinates(board game.Board) (byte, byte, error) {
	var err error = nil

	var y byte
	err = InputByte("Row number (0-2):", &y)

	if err != nil {
		return 0, 0, err
	}

	var x byte
	err = InputByte("Column number (0-2): ", &x)

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
