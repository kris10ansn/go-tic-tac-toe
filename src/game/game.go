package game

import "fmt"

const EMPTY_TIC Tic = ' '
const X_TIC Tic = 'x'
const O_TIC Tic = 'o'
const BOARD_WIDTH byte = 3
const BOARD_HEIGHT byte = 3

type Tic = byte
type Board = [BOARD_HEIGHT][BOARD_WIDTH]Tic

func CreateEmptyBoard() Board {
	return Board{
		{EMPTY_TIC, EMPTY_TIC, EMPTY_TIC},
		{EMPTY_TIC, EMPTY_TIC, EMPTY_TIC},
		{EMPTY_TIC, EMPTY_TIC, EMPTY_TIC},
	}
}

// TODO: search diagonals
func CheckWin(board Board) Tic {
	var vertical byte
	var horizontal byte

	for i, row := range board {
		for j, val := range row {
			if j == 0 {
				vertical = val
				horizontal = val
			}

			if board[i][j] != horizontal {
				horizontal = 0
			}

			if board[j][i] != vertical {
				vertical = 0
			}

			if vertical == 0 && horizontal == 0 {
				break
			}
		}

		if vertical != 0 {
			return vertical
		} else if horizontal != 0 {
			return horizontal
		}

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
			return fmt.Errorf("invalid turn: %c", *turn)
		}
	}
}

func BoardToString(board Board) string {
	str := ""

	for _, row := range board {
		for _, val := range row {
			str += fmt.Sprintf("[%c]", val)
		}
		str += "\n"
	}

	return str
}
