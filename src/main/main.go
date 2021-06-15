package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 5 x 6 game
type game struct {
	Board       [][]string
	PlayersTurn player
}

type player int

// TODO: Map a player to X or O

// printHelp diplays a helper message
func (g *game) printHelp() {
	fmt.Printf("\n%v\n", strings.Repeat("-", len(g.Board[0])*4+1))
	out := "|"
	for i := range g.Board[0] {
		out += " "
		out += strconv.Itoa(i)
		out += " |"
	}
	fmt.Println(out)
	g.printBoard()
}

// alternatePlayersTurn alternates the player whose move it is
func (g *game) alternatePlayersTurn() {
	if g.PlayersTurn == 0 {
		g.PlayersTurn = 1
		return
	}
	g.PlayersTurn = 0
}

// printBoard pretty prints the board
func (g *game) printBoard() {
	fmt.Println(strings.Repeat("-", len(g.Board[0])*4+1))
	for i, _ := range g.Board {
		out := "|"
		for j, _ := range g.Board[i] {
			out += " "
			out += g.Board[i][j]
			out += " |"
		}
		fmt.Println(out)
		fmt.Println(strings.Repeat("-", len(g.Board[0])*4+1))
	}
}

// init initializes the board and sets all fields to " "
func (g *game) init() {
	g.PlayersTurn = 0
	g.Board = make([][]string, 6)
	for i, _ := range g.Board {
		g.Board[i] = make([]string, 5)
		for j, _ := range g.Board[i] {
			g.Board[i][j] = " "
		}
	}
}

// selectComputerMove selects a random number
func (g *game) selectComputerMove() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6)
}

// selectHumanMove gets a move from the user
func (g *game) selectHumanMove() int {
	// TODO
	return 0
}

// selectMove wraps doComputerMove and doHumanMove
func (g *game) selectMove() int {
	if g.PlayersTurn == 0 {
		return g.selectHumanMove()
	}
	return g.selectComputerMove()
}

// TODO
func (g *game) won() (bool, player) {
	return false, -1
}

func main() {
	fmt.Printf("Let's play 4 wins!\nThe top line indicates the rows you can choose")
	var g game
	g.init()
	g.printHelp()

	// game loop
	gameFinished := false
	for !gameFinished {
		// players do action
		g.selectMove()
		gameFinished, _ = g.won()
	}
}
