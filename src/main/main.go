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

type game struct {
	Board       [][]string
	PlayersTurn player
	Width       int
	Height      int
}

type player int

type direction int

const (
	Horizontal    direction = iota
	Vertical                = iota
	Diagonal                = iota
	noPlayerValue           = " "
)

var playerMapping = [...]string{"X", "O"}

// playerToString provides mapping for player to string
func playerToString(p player) string {
	return playerMapping[p]
}

// playerIntToStrig is a wrapper for playerToString
func playerIntToStrig(p int) string {
	return playerToString(player(p))
}

// TODO: Remove this function?
// stringToPlayer provides mapping from a string to a player
func stringToPlayer(s string) player {
	for i, val := range playerMapping {
		if val == s {
			return player(i)
		}
	}
	panic("Unsupported player")
}

// TODO: Move Board struct and methods to other file

// printHelp diplays a helper message
func (g *game) printHelp() {
	fmt.Printf("\n%v\n", strings.Repeat("-", len(g.Board[0])*4+1))
	out := "|"
	for i := range g.Board[0] {
		out += noPlayerValue
		out += strconv.Itoa(i)
		out += " |"
	}
	fmt.Println(out)
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
			out += noPlayerValue
			out += g.Board[i][j]
			out += " |"
		}
		fmt.Println(out)
		fmt.Println(strings.Repeat("-", len(g.Board[0])*4+1))
	}
}

// init initializes the board and sets all fields to noPlayerValue
func (g *game) init() {
	g.PlayersTurn = 1
	g.Width = 6
	g.Height = 8
	g.Board = make([][]string, g.Height)
	for i, _ := range g.Board {
		g.Board[i] = make([]string, g.Width)
		for j, _ := range g.Board[i] {
			g.Board[i][j] = noPlayerValue
		}
	}
}

func (g *game) _selectComputerMove() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(g.Width)
}

// selectComputerMove is a wraper for _selectComputerMove
func (g *game) selectComputerMove() int {
	fmt.Println("Computer move")
	time.Sleep(1 * time.Second)
	return g._selectComputerMove()
}

// selectHumanMove gets a move from the user
func (g *game) selectHumanMove(reader *bufio.Reader) (int, error) {
	fmt.Println("Please select a column")
	var userInput int
	_, err := fmt.Scanf("%d", &userInput)
	if err != nil {
		return -1, err
	}
	if userInput < 0 || userInput >= g.Width {
		return -1, fmt.Errorf("please select a number between 0 and %v inclusive", g.Width-1)
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

// fall simulates the falling of the token in the column
// It returns the row in which the token will rest or
// an error if this column is already full
func (g *game) fall(column int) (int, error) {
	if g.Board[0][column] != noPlayerValue {
		return 0, fmt.Errorf("column %v is already full", column)
	}
	for row := 0; row < len(g.Board); row++ {
		if g.Board[row][column] != noPlayerValue {
			return row - 1, nil
		}
	}
	return len(g.Board) - 1, nil
}

// doMove does executes a selected move on the board
// returns an error if the selected column is already full
func (g *game) doMove(column int) error {
	row, err := g.fall(column)
	if err != nil {
		fmt.Printf("Column %v already full. Choose again\n", column)
		return err
	}
	g.Board[row][column] = playerToString(g.PlayersTurn)
	return nil
}

// checkKHorizontal checks if the next k horizontal to the right
// fields are equal to playerString
// Returns false, if out of bound
// If xPos is 0, the fields 0, 1, 2, and 3 will be looked at
func (g *game) checkKHorizontal(playerSting string, k int, xPos int, yPos int) bool {
	// if out of bound return false
	if xPos+k-1 >= g.Width {
		return false
	}

	for i := 0; i < k; i++ {
		if g.Board[yPos][xPos+i] != playerSting {
			return false
		}
	}
	return true
}

// checkKVertical checks if the next k horizontal to the bottom
// fields are equal to playerString
// Returns false, if out of bound
// If yPos is 0, the fields 0, 1, 2, and 3 will be looked at
func (g *game) checkKVertical(playerSting string, k int, xPos int, yPos int) bool {
	// if out of bound return false
	if yPos+k-1 >= g.Height {
		return false
	}

	for i := 0; i < k; i++ {
		if g.Board[yPos+i][xPos] != playerSting {
			return false
		}
	}
	return true
}

// TODO: Test diagonal winnings

// checkKDiagonalDown checks if the next k diagonal to the bottom right
// fields are equal to playerString
// Returns false, if out of bound
// If yPos and xPos are both 0, the fields 1,1, 2,2, and 3,3 will be looked at
func (g *game) checkKDiagonalDown(playerSting string, k int, xPos int, yPos int) bool {
	// if out of bound return false
	if yPos+k-1 >= g.Height || xPos+k-1 >= g.Width {
		return false
	}

	for i := 0; i < k; i++ {
		if g.Board[yPos+i][xPos+i] != playerSting {
			return false
		}
	}
	return true
}

// checkKDiagonalUp checks if the next k diagonal to the top right
// fields are equal to playerString
// Returns false, if out of bound
// If yPos and xPos are both 0, the fields 1,1, 2,2, and 3,3 will be looked at
func (g *game) checkKDiagonalUp(playerSting string, k int, xPos int, yPos int) bool {
	// if out of bound return false
	if yPos-k+1 < 0 || xPos+k-1 >= g.Width {
		return false
	}

	for i := 0; i < k; i++ {
		if g.Board[yPos-i][xPos+i] != playerSting {
			return false
		}
	}
	return true
}

// checkX checks if the next k in a given direction
// fields are equal to playerString
// Returns false, if out of bound
func (g *game) checkX(playerString string, k int, xPos int, yPos int, d direction) bool {
	if d == Horizontal {
		return g.checkKHorizontal(playerString, k, xPos, yPos)
	} else if d == Vertical {
		return g.checkKVertical(playerString, k, xPos, yPos)
	} else if d == Diagonal {
		return g.checkKDiagonalDown(playerString, k, xPos, yPos) ||
			g.checkKDiagonalUp(playerString, k, xPos, yPos)
	}
	panic(fmt.Sprintf("Unsupported direction %v", d))
}

// won returns the winning player and true
// if no player won, it returns false
func (g *game) won() (bool, player) {
	for yPos := 0; yPos < g.Height; yPos++ {
		for xPos := 0; xPos < g.Width; xPos++ {
			// TODO: Map player to strings here

			// horizontal
			for i := 0; i <= 1; i++ {
				if g.checkX(playerIntToStrig(i), 4, xPos, yPos, Horizontal) {
					return true, player(i)
				}
			}

			// vertical
			for i := 0; i <= 1; i++ {
				if g.checkX(playerIntToStrig(i), 4, xPos, yPos, Vertical) {
					return true, player(i)
				}
			}

			// diagonal
			for i := 0; i <= 1; i++ {
				if g.checkX(playerIntToStrig(i), 4, xPos, yPos, Diagonal) {
					return true, player(i)
				}
			}
		}
	}

	return false, -1
}

// boardFull returns true if no more moves are possible
func (g *game) boardFull() bool {
	for _, val := range g.Board[0] {
		if val == noPlayerValue {
			return false
		}
	}
	return true
}

func getNewGame() game {
	var g game
	g.init()
	return g
}

func main() {
	fmt.Printf("Let's play 4 wins!\nThe top line indicates the rows you can choose\n")
	g := getNewGame()

	ioReader := bufio.NewReader(os.Stdin)

	// game loop
	for !g.boardFull() {
		// players do action
		columnSelected := g.selectMove(ioReader)
		err := g.doMove(columnSelected)
		if err != nil {
			continue
		}
		g.printHelp()
		g.printBoard()
		gameFinished, winner := g.won()
		if gameFinished {
			fmt.Printf("Congrats to player %v. You are a winner!\n", playerToString(winner))
			return
		}
		g.alternatePlayersTurn()
	}
	fmt.Println("Board full. Game ends")
}
