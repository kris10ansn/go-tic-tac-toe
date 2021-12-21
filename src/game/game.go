package game

import "fmt"

const EMPTY_TIC Tic = ' '
const X_TIC Tic = 'x'
const O_TIC Tic = 'o'
const BOARD_WIDTH byte = 3
const BOARD_HEIGHT byte = 3

type Tic = byte
type Board = [3][3]Tic

func CreateEmptyBoard() Board {
	return Board{
		{EMPTY_TIC, EMPTY_TIC, EMPTY_TIC},
		{EMPTY_TIC, EMPTY_TIC, EMPTY_TIC},
		{EMPTY_TIC, EMPTY_TIC, EMPTY_TIC},
	}
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
