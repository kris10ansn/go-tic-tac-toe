package cli

import (
	"fmt"

	"github.com/kris10ansn/go-tic-tac-toe/pkg/game"
)

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
