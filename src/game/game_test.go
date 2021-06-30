package game

import (
	"fmt"
	"fourwins/main/src/misc"
	"testing"
)

func TestUsableBoard(t *testing.T) {
	g := GetNewGame()
	if g.Width < 4 {
		t.Fatalf("Board has incorrect with of %v. Board with needs to be at least 4", g.Width)
	}
}

func TestCheckNextXHorizontal(t *testing.T) {
	g := GetNewGame()

	for p := 0; p <= 1; p++ {
		for i := 0; i < 4; i++ {
			g.Board[0][i] = misc.PlayerIntToStrig(p)
		}
		won, winingPlayer := g.Won()
		if !won {
			g.PrintBoard()
			t.Error("expected game to end because player won")
		} else if winingPlayer != misc.Player(p) {
			g.PrintBoard()
			t.Errorf("expected player %v to win", misc.PlayerIntToStrig(p))
		}
	}

	g.init()
}

func Test_selectComputerMove(t *testing.T) {
	var g game
	g.init()
	for i := 0; i < 100; i++ {
		selectedMove := g._selectComputerMove()
		if selectedMove < 0 || selectedMove >= g.Width {
			t.Errorf("Expected selected move to be [0, g.Width)")
		}
	}
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
		{[]int{0, 0, 0, 0}, []int{0, 1, 2, 3}},
		{[]int{0, 0, 0, 0}, []int{1, 2, 3, 4}},
		{[]int{0, 0, 0, 0}, []int{2, 3, 4, 5}},
		{[]int{0, 0, 0, 0}, []int{g.Height - 1, g.Height - 2, g.Height - 3, g.Height - 4}},
		{[]int{g.Width - 1, g.Width - 1, g.Width - 1, g.Width - 1}, []int{0, 1, 2, 3}},
		{[]int{g.Width - 1, g.Width - 1, g.Width - 1, g.Width - 1}, []int{1, 2, 3, 4}},
		{[]int{g.Width - 1, g.Width - 1, g.Width - 1, g.Width - 1}, []int{2, 3, 4, 5}},
		{[]int{g.Width - 1, g.Width - 1, g.Width - 1, g.Width - 1}, []int{g.Height - 1, g.Height - 2, g.Height - 3, g.Height - 4}},

		// diagonal missing
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{2, 3, 4, 5}, []int{2, 3, 4, 5}},
		{[]int{5, 4, 3, 2}, []int{2, 3, 4, 5}},
		{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
	}

	for _, scenario := range scenarios {
		for i, xpos := range scenario.xPos {
			g.Board[scenario.yPos[i]][xpos] = misc.PlayerIntToStrig(0)
		}
		won, winning_player := g.Won()
		if !won {
			g.PrintBoard()
			msg := "Expected a player to have one, but was not the case. See above for game filed"
			fmt.Println(msg)
			t.Errorf(msg)
		}
		if won && int(winning_player) != 0 {
			g.PrintBoard()
			msg := "Expected a player 0 to have one, but was not the case. See above for game filed"
			fmt.Println(msg)
			t.Errorf(msg)
		}

		g.init()
	}
}