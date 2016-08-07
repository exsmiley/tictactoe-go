package player

import (
	"testing"
)

func TestNextMove(t *testing.T) {
	board := [][]string{{"", "x", ""},{"", "o", "x"},{"o", "", "x"}}
	move := GetNextMove(board)
	if move.X != 0 || move.Y != 2 {
		t.Error("Expected", 0, 2, "Got", move)
	}

	board = [][]string{{"o", "x", ""},{"", "x", "o"},{"", "", ""}}
	move = GetNextMove(board)
	if move.X != 2 || move.Y != 1 {
		t.Error("Expected", 2, 1, "Got", move)
	}

	board = [][]string{{"o", "x", "o"},{"", "x", "o"},{"", "", ""}}
	move = GetNextMove(board)
	if move.X != 2 || move.Y != 2 {
		t.Error("Expected", 2, 2, "Got", move)
	}

	board = [][]string{{"o", "", ""},{"", "", ""},{"", "x", "x"}}
	move = GetNextMove(board)
	if move.X != 2 || move.Y != 0 {
		t.Error("Expected", 2, 0, "Got", move)
	}

	board = [][]string{{"", "", "x"},{"", "", ""},{"", "o", "x"}}
	move = GetNextMove(board)
	if move.X != 1 || move.Y != 2 {
		t.Error("Expected", 1, 2, "Got", move)
	}

	board = [][]string{{"o", "", "x"},{"", "x", ""},{"", "", ""}}
	move = GetNextMove(board)
	if move.X != 2 || move.Y != 0 {
		t.Error("Expected", 2, 0, "Got", move)
	}
	board = [][]string{{"o", "", "x"},{"", "", ""},{"x", "", ""}}
	move = GetNextMove(board)
	if move.X != 1 || move.Y != 1 {
		t.Error("Expected", 1, 1, "Got", move)
	}
	board = [][]string{{"x", "", ""},{"", "", ""},{"", "", ""}}
	move = GetNextMove(board)
	if move.X != 1 || move.Y != 1 {
		t.Error("Expected", 1, 1, "Got", move)
	}
}

type gameovertest struct {
	board [][]string
	over bool
	winner string
}

var gameOverTests = []gameovertest {
	{[][]string{{"o", "o", "o"},{"o", "o", "o"},{"o", "x", "x"}}, true, "o"},
	{[][]string{{"", "", ""},{"", "", ""},{"", "", ""}}, false, ""},
	{[][]string{{"", "", ""},{"x", "x", "x"},{"", "", ""}}, true, "x"},
	{[][]string{{"o", "x", "x"},{"", "o", ""},{"", "", "o"}}, true, "o"},
	{[][]string{{"x", "x", "x"},{"", "o", ""},{"", "", "o"}}, true, "x"},
	{[][]string{{"o", "x", "o"},{"x", "o", ""},{"o", "x", "o"}}, true, "o"},
	{[][]string{{"x", "o", "x"},{"", "x", "o"},{"o", "", "x"}}, true, "x"},
}

func TestIsGameOver(t *testing.T) {
	for _, test := range gameOverTests {
		board := test.board
		over := test.over
		winner := test.winner
		ansOver, ansWinner := IsGameOver(board)
		if over != ansOver || winner != ansWinner {
			t.Error("For", board, "expected", over, winner, "got", ansOver, ansWinner)
		}
	}
}

func TestGetNextPlayer(t *testing.T) {
	player := "x"
	next := getNextPlayer(player)
	if next != "o" {
		t.Error("Expected", "o", "Got", next)
	}

	player = "o"
	next = getNextPlayer(player)
	if next != "x" {
		t.Error("Expected", "x", "Got", next)
	}
}

func TestIsEmptyBoard(t *testing.T) {
	blankBoard := [][]string{{"","",""},{"","",""},{"","",""}}

	blank := isEmpty(blankBoard)

	if !blank {
		t.Error("Expected true, got", blank)
	}

	notBlankBoard := [][]string{{"","",""},{"x","",""},{"","",""}}

	notBlank := isEmpty(notBlankBoard)

	if notBlank {
		t.Error("Expected false, got", notBlank)
	}
}