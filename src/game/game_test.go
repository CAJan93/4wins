package game

import (
	"fmt"
	"fourwins/main/src/misc"
	"testing"
)

// Try to catch the error that shows up when first slecting an invalid column
// Please select a column
// 9
// Invalid input, please try again
// Please select a column
// 3
// panic: runtime error: index out of range [-1]
//
// goroutine 1 [running]:
// fourwins/main/src/game.(*Game).fall(0xc0000a1e98, 0xffffffffffffffff, 0x10b8d16, 0xc0000a1e98, 0xc0000a1f20)
//         /Users/D072532/go/src/4Gewinnt/src/game/game.go:242 +0x16e
// fourwins/main/src/game.(*Game).DoMove(0xc0000a1e98, 0xffffffffffffffff, 0xffffffffffffffff, 0x0)
//         /Users/D072532/go/src/4Gewinnt/src/game/game.go:256 +0x3f
// main.main()
//         /Users/D072532/go/src/4Gewinnt/src/main/main.go:21 +0x24f
// make: *** [run] Error 2

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

// Tests obvious decisions, where 3 objects are in a row already
func Test_selectComputerMoveDecision(t *testing.T) {
	type testHelper struct {
		xPos         []int
		yPos         []int
		expectedMove int
	}
	var g Game
	g.init()
	bottom := g.Height - 1

	scenarios := [...]testHelper{
		// horizontal
		{[]int{0, 1, 2}, []int{bottom, bottom, bottom}, 3},
		{[]int{g.Width - 1, g.Width - 2, g.Width - 3}, []int{bottom, bottom, bottom}, g.Width - 4},

		// vertical
		{[]int{0, 0, 0}, []int{bottom, bottom - 1, bottom - 2}, 0},
		{[]int{g.Width - 1, g.Width - 1, g.Width - 1}, []int{bottom, bottom - 1, bottom - 2}, g.Width - 1},
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

func TestWon(t *testing.T) {
	type testHelper struct {
		xPos []int
		yPos []int
	}

	var g Game
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
		{[]int{0, 1, 2, 3, 4, 5}, []int{g.Height - 1, g.Height - 1, g.Height - 1, g.Height - 1, g.Height - 1, g.Height - 1}},

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
