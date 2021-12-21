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
	for _, row := range board {
		for _, val := range row {
			fmt.Printf("[%c]", val)
		}
		fmt.Println()
	}
}

func InputCoordinates(gameBoard game.Board) (byte, byte, error) {
	var err error = nil

	var y byte
	err = InputByte("Row number:", &y)

	if err != nil {
		return 0, 0, err
	}

	var x byte
	err = InputByte("Column number: ", &x)

	if err != nil {
		return 0, 0, err
	}

	if x >= game.BOARD_WIDTH || y >= game.BOARD_HEIGHT {
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
