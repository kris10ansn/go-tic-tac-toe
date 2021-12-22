package cli

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/src/game"
)

func InputByte(message string, input *byte) error {
	fmt.Print(message)
	_, err := fmt.Scanf("%d\n", input)

	return err
}

func PrintBoard(board game.Board) {
	fmt.Println(game.BoardToString(board))
}

func PrintGameEnd(board game.Board, winner game.Tic, moves byte) {
	if winner != game.EMPTY_TIC {
		fmt.Printf("%s won the game after %d moves!\n", game.TicToString(winner), moves)
	} else if moves == 9 {
		fmt.Println("It's a draw!")
	} else {
		panic("Unexpected behaviour")
	}
}

func InputCoordinates(gameBoard game.Board) (byte, byte, error) {
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

	if gameBoard[y][x] != game.EMPTY_TIC {
		return x, y, fmt.Errorf("[%d, %d] is already occupied", x, y)
	}

	return x, y, nil
}

func WaitForMove(gameBoard game.Board) (byte, byte) {
	var x byte
	var y byte
	var err error

	for x, y, err = InputCoordinates(gameBoard); err != nil; {
		fmt.Printf("%s\n", err)
		x, y, err = InputCoordinates(gameBoard)
	}

	return x, y
}
