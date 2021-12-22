package game

import (
	"testing"
)

type Row = [3]Tic

var tics = [2]Tic{X_TIC, O_TIC}
var empty_row = Row{EMPTY_TIC, EMPTY_TIC, EMPTY_TIC}
var horizontal_x_row = Row{X_TIC, X_TIC, X_TIC}
var horizontal_o_row = Row{O_TIC, O_TIC, O_TIC}

type CheckWinTest struct {
	input    Board
	expected Tic
}

func TestBoardToString(t *testing.T) {
	board := CreateEmptyBoard()
	board_string := BoardToString(board)
	board_string_expected := "[ ][ ][ ]\n[ ][ ][ ]\n[ ][ ][ ]"

	if board_string != board_string_expected {
		t.Errorf(
			"Test failed: empty board inputted, \n\"%s\" expected, \n\"%s\" recieved",
			board_string_expected,
			board_string,
		)
	}
}

func TestTableTicToString(t *testing.T) {
	var tests = []struct {
		input    Tic
		expected string
	}{
		{X_TIC, "x"},
		{O_TIC, "o"},
		{EMPTY_TIC, " "},
	}

	for i, test := range tests {
		if output := TicToString(test.input); output != test.expected {
			t.Errorf(
				"Test %d failed: %d inputted, \"%s\" expected, \"%s\" recieved",
				i,
				test.input,
				test.expected,
				output,
			)
		}
	}
}

func TestTableNextTurn(t *testing.T) {
	var prevTic Tic = X_TIC
	var tic Tic = prevTic
	var expected Tic = O_TIC

	err := NextTurn(&tic)

	if tic != expected {
		t.Errorf(
			"Test 0 failed: %s => NextTurn => %s, expected %s",
			TicToString(prevTic),
			TicToString(tic),
			TicToString(expected),
		)
	}

	if err != nil {
		t.Errorf("Test 0 error: %s", err)
	}

	err = NextTurn(&tic)
	expected = X_TIC

	if tic != expected {
		t.Errorf(
			"Test 1 failed: %s => NextTurn => %s, expected %s",
			TicToString(prevTic),
			TicToString(tic),
			TicToString(expected),
		)
	}

	if err != nil {
		t.Errorf("Test 1 error: %s", err)
	}

	var emptyTic Tic = EMPTY_TIC
	err = NextTurn(&emptyTic)

	if err == nil {
		t.Error("Test 2 failed: Should have errored, no error recieved (NextTurn(EMPTY_TIC))")
	}
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

		// Non-winning board that could break CheckWin if X_TIC and O_TIC are not prime
		{Board{empty_row, empty_row, {EMPTY_TIC, O_TIC, X_TIC}}, EMPTY_TIC},

		// Randomly thought of no winner boards
		{Board{{O_TIC, X_TIC, O_TIC}, {X_TIC, O_TIC, X_TIC}, {X_TIC, O_TIC, X_TIC}}, EMPTY_TIC},
		{Board{{EMPTY_TIC, X_TIC, O_TIC}, {X_TIC, EMPTY_TIC, X_TIC}, {X_TIC, O_TIC, X_TIC}}, EMPTY_TIC},
		{Board{{EMPTY_TIC, X_TIC, O_TIC}, {X_TIC, EMPTY_TIC, X_TIC}, empty_row}, EMPTY_TIC},
		{Board{empty_row, {X_TIC, EMPTY_TIC, X_TIC}, empty_row}, EMPTY_TIC},
	}

	for i, test := range tests {
		if output := CheckWin(test.input); output != test.expected {
			t.Errorf(
				"\nTest %d failed: \n%s inputted, \n%s/%d expected, \n%s/%d recieved",
				i,
				BoardToString(test.input),
				TicToString(test.expected),
				test.expected,
				TicToString(output),
				output,
			)
		}
	}
}
