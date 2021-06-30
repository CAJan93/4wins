package main

import (
	"fmt"
	"testing"
)

func TestUsableBoard(t *testing.T) {
	g := getNewGame()
	if g.Width < 4 {
		t.Fatalf("Board has incorrect with of %v. Board with needs to be at least 4", g.Width)
	}
}

func TestCheckNextXHorizontal(t *testing.T) {
	g := getNewGame()

	for p := 0; p <= 1; p++ {
		for i := 0; i < 4; i++ {
			g.Board[0][i] = playerIntToStrig(p)
		}
		won, winingPlayer := g.won()
		if !won {
			g.printBoard()
			t.Error("expected game to end because player won")
		} else if winingPlayer != player(p) {
			g.printBoard()
			t.Errorf("expected player %v to win", playerIntToStrig(p))
		}
	}

	g.init()
}

func TestWon(t *testing.T) {
	type testHelper struct {
		xPos []int
		yPos []int
	}

	var g game
	g.init()

	scenarios := [...]testHelper{
		// horizontal
		{[]int{0, 1, 2, 3}, []int{0, 0, 0, 0}},
		{[]int{1, 2, 3, 4}, []int{0, 0, 0, 0}},
		{[]int{2, 3, 4, 5}, []int{0, 0, 0, 0}},
		{[]int{0, 1, 2, 3}, []int{1, 1, 1, 1}},
		{[]int{1, 2, 3, 4}, []int{1, 1, 1, 1}},
		{[]int{2, 3, 4, 5}, []int{1, 1, 1, 1}},
		{[]int{0, 1, 2, 3}, []int{2, 2, 2, 2}},
		{[]int{1, 2, 3, 4}, []int{2, 2, 2, 2}},
		{[]int{2, 3, 4, 5}, []int{2, 2, 2, 2}},
		{[]int{0, 1, 2, 3}, []int{g.Height - 1, g.Height - 1, g.Height - 1, g.Height - 1}},
		{[]int{1, 2, 3, 4}, []int{g.Height - 1, g.Height - 1, g.Height - 1, g.Height - 1}},
		{[]int{2, 3, 4, 5}, []int{g.Height - 1, g.Height - 1, g.Height - 1, g.Height - 1}},

		// vertical

		// diagonal missing
	}

	for _, scenario := range scenarios {
		for i, xpos := range scenario.xPos {
			g.Board[scenario.yPos[i]][xpos] = playerIntToStrig(0)
		}
		won, winning_player := g.won()
		if !won {
			g.printBoard()
			msg := "Expected a player to have one, but was not the case. See above for game filed"
			fmt.Println(msg)
			t.Errorf(msg)
		}
		if won && int(winning_player) != 0 {
			g.printBoard()
			msg := "Expected a player 0 to have one, but was not the case. See above for game filed"
			fmt.Println(msg)
			t.Errorf(msg)
		}

		g.init()
	}
}
