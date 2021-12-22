package game

import (
	"testing"
)

type Row = [BOARD_WIDTH]Tic

var tics = [2]Tic{X_TIC, O_TIC}
var empty_row = Row{EMPTY_TIC, EMPTY_TIC, EMPTY_TIC}
var horizontal_x_row = Row{X_TIC, X_TIC, X_TIC}
var horizontal_o_row = Row{O_TIC, O_TIC, O_TIC}

type CheckWinTest struct {
	input    Board
	expected Tic
}

func TestTableCheckWin(t *testing.T) {
	var tests = []CheckWinTest{
		// Horizontal x boards
		{Board{horizontal_x_row, empty_row, empty_row}, X_TIC},
		{Board{empty_row, horizontal_x_row, empty_row}, X_TIC},
		{Board{empty_row, empty_row, horizontal_x_row}, X_TIC},
		// Horizontal o boards
		{Board{horizontal_o_row, empty_row, empty_row}, O_TIC},
		{Board{empty_row, horizontal_o_row, empty_row}, O_TIC},
		{Board{empty_row, empty_row, horizontal_o_row}, O_TIC},

		// Vertical x boards
		{Board{{X_TIC, EMPTY_TIC, EMPTY_TIC}, {X_TIC, EMPTY_TIC, EMPTY_TIC}, {X_TIC, EMPTY_TIC, EMPTY_TIC}}, X_TIC},
		{Board{{EMPTY_TIC, X_TIC, EMPTY_TIC}, {EMPTY_TIC, X_TIC, EMPTY_TIC}, {EMPTY_TIC, X_TIC, EMPTY_TIC}}, X_TIC},
		{Board{{EMPTY_TIC, EMPTY_TIC, X_TIC}, {EMPTY_TIC, EMPTY_TIC, X_TIC}, {EMPTY_TIC, EMPTY_TIC, X_TIC}}, X_TIC},

		// Vertical o boards
		{Board{{O_TIC, EMPTY_TIC, EMPTY_TIC}, {O_TIC, EMPTY_TIC, EMPTY_TIC}, {O_TIC, EMPTY_TIC, EMPTY_TIC}}, O_TIC},
		{Board{{EMPTY_TIC, O_TIC, EMPTY_TIC}, {EMPTY_TIC, O_TIC, EMPTY_TIC}, {EMPTY_TIC, O_TIC, EMPTY_TIC}}, O_TIC},
		{Board{{EMPTY_TIC, EMPTY_TIC, O_TIC}, {EMPTY_TIC, EMPTY_TIC, O_TIC}, {EMPTY_TIC, EMPTY_TIC, O_TIC}}, O_TIC},

		// Diagonal x boards
		{Board{{X_TIC, EMPTY_TIC, EMPTY_TIC}, {EMPTY_TIC, X_TIC, EMPTY_TIC}, {EMPTY_TIC, EMPTY_TIC, X_TIC}}, X_TIC},
		{Board{{EMPTY_TIC, EMPTY_TIC, X_TIC}, {EMPTY_TIC, X_TIC, EMPTY_TIC}, {X_TIC, EMPTY_TIC, EMPTY_TIC}}, X_TIC},

		// Diagonal o boards
		{Board{{O_TIC, EMPTY_TIC, EMPTY_TIC}, {EMPTY_TIC, O_TIC, EMPTY_TIC}, {EMPTY_TIC, EMPTY_TIC, O_TIC}}, O_TIC},
		{Board{{EMPTY_TIC, EMPTY_TIC, O_TIC}, {EMPTY_TIC, O_TIC, EMPTY_TIC}, {O_TIC, EMPTY_TIC, EMPTY_TIC}}, O_TIC},
	}

	for _, test := range tests {
		if output := CheckWin(test.input); output != test.expected {
			t.Errorf(
				"\nTest Failed: \n%s inputted, \n%s expected, \n%s recieved",
				BoardToString(test.input),
				TicToString(test.expected),
				TicToString(output),
			)
		}
	}
}
