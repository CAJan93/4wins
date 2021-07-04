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

func Test_selectComputerMoveInBound(t *testing.T) {
	var g Game
	g.init()
	for i := 0; i < 5; i++ {
		selectedMove := g._selectComputerMove()
		if selectedMove < 0 || selectedMove >= g.Width {
			t.Errorf("Expected selected move to be [0, g.Width). Selected move was %v", selectedMove)
		}
		g.init()
	}
}

func Benchmark_minmax(b *testing.B) {
	var g Game
	g.init()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		minmax(g, 5, misc.Player(0))
	}
}

// Tests obvious decisions, where 3 objects are in a row already
func Test_selectComputerMoveDecision(t *testing.T) {
	type testHelper struct {
		xPos         []int
		yPos         []int
		expectedMove int
	}
	var g Game
	g.init()
	bottom := gHeight() - 1

	scenarios := [...]testHelper{
		// horizontal
		{[]int{0, 1, 2}, []int{bottom, bottom, bottom}, 3},
		{[]int{gWidth() - 1, g.Width - 2, g.Width - 3}, []int{bottom, bottom, bottom}, g.Width - 4},

		// vertical
		{[]int{0, 0, 0}, []int{bottom, bottom - 1, bottom - 2}, 0},
		{[]int{gWidth() - 1, gWidth() - 1, gWidth() - 1}, []int{bottom, bottom - 1, bottom - 2}, gWidth() - 1},
	}

	// do for both players
	for player := 0; player <= 1; player++ {
		for _, scenario := range scenarios {
			for i := 0; i < len(scenario.xPos); i++ {
				g.Board[scenario.yPos[i]][scenario.xPos[i]] = misc.PlayerIntToStrig(player)
			}
			move := g._selectComputerMove()
			if move != scenario.expectedMove {
				msg := fmt.Sprintf("Expected move %v, got move %v", scenario.expectedMove, move)
				g.PrintBoard()
				fmt.Println(msg)
				t.Errorf(msg)
			}
			g.init()
		}
	}
}

type testHelper struct {
	xPos []int
	yPos []int
}

func gHeight() int {
	var g Game
	g.init()
	return g.Height
}

func gWidth() int {
	var g Game
	g.init()
	return g.Width
}

var scenariosTestWon = [...]testHelper{
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
	{[]int{0, 1, 2, 3}, []int{gHeight() - 1, gHeight() - 1, gHeight() - 1, gHeight() - 1}},
	{[]int{1, 2, 3, 4}, []int{gHeight() - 1, gHeight() - 1, gHeight() - 1, gHeight() - 1}},
	{[]int{2, 3, 4, 5}, []int{gHeight() - 1, gHeight() - 1, gHeight() - 1, gHeight() - 1}},
	{[]int{0, 1, 2, 3, 4, 5}, []int{gHeight() - 1, gHeight() - 1, gHeight() - 1, gHeight() - 1, gHeight() - 1, gHeight() - 1}},

	// vertical
	{[]int{0, 0, 0, 0}, []int{0, 1, 2, 3}},
	{[]int{0, 0, 0, 0}, []int{1, 2, 3, 4}},
	{[]int{0, 0, 0, 0}, []int{2, 3, 4, 5}},
	{[]int{0, 0, 0, 0}, []int{gHeight() - 1, gHeight() - 2, gHeight() - 3, gHeight() - 4}},
	{[]int{gWidth() - 1, gWidth() - 1, gWidth() - 1, gWidth() - 1}, []int{0, 1, 2, 3}},
	{[]int{gWidth() - 1, gWidth() - 1, gWidth() - 1, gWidth() - 1}, []int{1, 2, 3, 4}},
	{[]int{gWidth() - 1, gWidth() - 1, gWidth() - 1, gWidth() - 1}, []int{2, 3, 4, 5}},
	{[]int{gWidth() - 1, gWidth() - 1, gWidth() - 1, gWidth() - 1}, []int{gHeight() - 1, gHeight() - 2, gHeight() - 3, gHeight() - 4}},

	// diagonal missing
	{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	{[]int{2, 3, 4, 5}, []int{2, 3, 4, 5}},
	{[]int{5, 4, 3, 2}, []int{2, 3, 4, 5}},
	{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
}

func TestWon(t *testing.T) {

	var g Game
	g.init()

	for _, scenario := range scenariosTestWon {
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

func BenchmarkWon(b *testing.B) {
	var g Game
	g.init()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		g.Won()
	}
}
