package game

import (
	"fmt"
)

/*
	CheckWin depends on X_TIC and O_TIC being prime numbers, so that only
	a combination of three X's or O's can result in the sum of that line
	being three times the value of that tic.
*/
const (
	X_TIC     Tic = 11
	O_TIC     Tic = 13
	EMPTY_TIC Tic = 0
)

type (
	Tic   = byte
	Board = [3][3]Tic
)

func CreateEmptyBoard() *Board {
	return &Board{
		{EMPTY_TIC, EMPTY_TIC, EMPTY_TIC},
		{EMPTY_TIC, EMPTY_TIC, EMPTY_TIC},
		{EMPTY_TIC, EMPTY_TIC, EMPTY_TIC},
	}
}

func SetBoardCoordinate(board *Board, x int, y int, tic Tic) {
	board[y][x] = tic
}

func CheckWin(board *Board) Tic {
	const X_WIN = 3 * X_TIC
	const O_WIN = 3 * O_TIC

	var vertical byte
	var horizontal byte
	var diagonal1 byte
	var diagonal2 byte

	for i, row := range board {
		vertical = 0
		horizontal = 0

		diagonal1 += board[i][i]
		diagonal2 += board[i][2-i]

		for j := range row {
			horizontal += board[i][j]
			vertical += board[j][i]
		}

		if vertical == X_WIN || vertical == O_WIN {
			return vertical / 3
		}
		if horizontal == X_WIN || horizontal == O_WIN {
			return horizontal / 3
		}
	}

	if diagonal1 == X_WIN || diagonal1 == O_WIN {
		return diagonal1 / 3
	}
	if diagonal2 == X_WIN || diagonal2 == O_WIN {
		return diagonal2 / 3
	}

	return EMPTY_TIC
}

func NextTurn(turn *Tic) error {
	switch *turn {
	case X_TIC:
		{
			*turn = O_TIC
			return nil
		}
	case O_TIC:
		{
			*turn = X_TIC
			return nil
		}
	default:
		{
			return fmt.Errorf("invalid turn: %d", *turn)
		}
	}
}

func BoardToString(board *Board) string {
	str := ""

	for i := 0; i < 3; i++ {
		for _, val := range board[i] {
			str += fmt.Sprintf("[%s]", TicToString(val))
		}
		if i != 3-1 {
			str += "\n"
		}
	}

	return str
}

func TicToString(tic Tic) string {
	switch tic {
	case X_TIC:
		return "x"
	case O_TIC:
		return "o"
	default:
		return " "
	}
}
