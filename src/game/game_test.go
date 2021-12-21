package game

import (
	"testing"
)

type Row = [BOARD_WIDTH]Tic

var tics = [2]Tic{X_TIC, O_TIC}
var empty_row = Row{EMPTY_TIC, EMPTY_TIC, EMPTY_TIC}
var horizontal_x_row = Row{X_TIC, X_TIC, X_TIC}
var horizontal_o_row = Row{O_TIC, O_TIC, O_TIC}

func TestTableCheckWin(t *testing.T) {
	var tests = []struct {
		input    Board
		expected Tic
	}{
		// Horizontal x boards
		{Board{horizontal_x_row, empty_row, empty_row}, X_TIC},
		{Board{empty_row, horizontal_x_row, empty_row}, X_TIC},
		{Board{empty_row, empty_row, horizontal_x_row}, X_TIC},
		// Horizontal o boards
		{Board{horizontal_o_row, empty_row, empty_row}, O_TIC},
		{Board{empty_row, horizontal_o_row, empty_row}, O_TIC},
		{Board{empty_row, empty_row, horizontal_o_row}, O_TIC},

		// TODO: vertical tests, diagonal tests, generated tests?
	}

	for _, test := range tests {
		if output := CheckWin(test.input); output != test.expected {
			t.Errorf(
				"\nTest Failed: \n%s inputted, \n%c expected, \n%c recieved",
				BoardToString(test.input),
				test.expected,
				output,
			)
		}
	}
}
