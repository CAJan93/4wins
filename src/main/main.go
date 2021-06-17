package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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
	g.PlayersTurn = 1
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
	fmt.Println("Computer move")
	time.Sleep(1 * time.Second)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(5)
}

// selectHumanMove gets a move from the user
func (g *game) selectHumanMove(reader *bufio.Reader) (int, error) {
	fmt.Println("Please select a column")
	var userInput int
	_, err := fmt.Scanf("%d", &userInput)
	if err != nil {
		return -1, err
	}
	if userInput < 0 || userInput > 4 {
		return -1, fmt.Errorf("please select a number between 0 and 4 inclusive")
	}
	return userInput, nil
}

// selectMove wraps doComputerMove and doHumanMove
func (g *game) selectMove(reader *bufio.Reader) int {
	if g.PlayersTurn == 0 {
		num, err := g.selectHumanMove(reader)
		if err != nil {
			fmt.Println(err) // TODO remove line
			fmt.Println("Invalid input, please try again")
			g.selectMove(reader)
		}
		return num
	}
	return g.selectComputerMove()
}

// doMove does executes a selected move on the board
func (g *game) doMove(column int) {
	if g.PlayersTurn == 0 {
		g.Board[0][column] = "X"
		return
	}
	g.Board[0][column] = "O"
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

	ioReader := bufio.NewReader(os.Stdin)

	// game loop
	gameFinished := false
	for !gameFinished {
		// players do action
		columnSelected := g.selectMove(ioReader)
		g.doMove(columnSelected)
		g.printBoard()
		gameFinished, _ = g.won()
		g.alternatePlayersTurn()
	}
}
