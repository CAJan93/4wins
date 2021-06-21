package game

import "testing"

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
