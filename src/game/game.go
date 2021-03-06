package game

import (
	"bufio"
	"fmt"
	"fourwins/main/src/misc"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	MAXIMIZINGPLAYER misc.Player = 0
	MINIMIZINGPLAYER misc.Player = 1
)

// getAllPossiblePossitions returns all possible branches in a decision inside the minmax tree
func getAllPossiblePossitions(incomingGame Game, maximizingPlayer bool) ([]Game, []int) {
	var possibleGameStates []Game
	var possibleMoves []int

	for column := 0; column < incomingGame.Width; column++ {
		tmp := incomingGame.copyBoard()
		if maximizingPlayer {
			if tmp.PlayersTurn != MAXIMIZINGPLAYER {
				tmp.AlternatePlayersTurn()
			}
		} else {
			if tmp.PlayersTurn != MINIMIZINGPLAYER {
				tmp.AlternatePlayersTurn()
			}
		}
		err := tmp.DoMove(column)
		if err != nil {
			continue // full column
		}
		possibleGameStates = append(possibleGameStates, tmp)
		possibleMoves = append(possibleMoves, column)
	}

	return possibleGameStates, possibleMoves
}

// _minmax function, compare https://www.youtube.com/watch?v=l-hh51ncgDI
func _minmax(position Game, remainingDepth float64, maximizingPlayer bool, lastMove int) (float64, int) {
	// handle winning position
	won, winningPlayer := position.Won()
	if won {
		if winningPlayer == MAXIMIZINGPLAYER {
			return 1, lastMove
		}
		return -1, lastMove
	}

	// no more search space or no more game space available
	if remainingDepth == 0 || position.BoardFull() {
		return 0, lastMove
	}

	if maximizingPlayer {
		maxEval := math.Inf(-1)
		bestMove := -1
		gameStates, gameMoves := getAllPossiblePossitions(position, maximizingPlayer)
		n := 2
		for i, child := range gameStates {
			eval, _ := _minmax(child, remainingDepth-1, false, gameMoves[i])
			if eval == 1 { // stop early on optimal move
				return 1, gameMoves[i]
			}
			if eval > maxEval || (eval == maxEval && rand.Intn(10) > 3) {
				maxEval = eval
				bestMove = gameMoves[i]
				n = 2
			} else if eval == maxEval { // pick equally good move by chance
				if rand.Intn(n) == 1 {
					bestMove = gameMoves[i]
				}
				n++
			}
		}
		return maxEval, bestMove
	}

	minEval := math.Inf(1)
	bestMove := -1
	gameStates, gameMoves := getAllPossiblePossitions(position, maximizingPlayer)
	n := 2
	for i, child := range gameStates {
		eval, _ := _minmax(child, remainingDepth-1, true, gameMoves[i])
		if eval == -1 {
			return -1, gameMoves[i]
		}
		if eval < minEval || (eval == minEval && rand.Intn(10) > 3) {
			minEval = eval
			bestMove = gameMoves[i]
		} else if eval == minEval { // pick equally good move by chance
			if rand.Intn(n) == 1 {
				bestMove = gameMoves[i]
			}
			n++
		}
	}
	return minEval, bestMove
}

// minmax is a wrapper for _minmax
func minmax(position Game, depth float64, currentPlayer misc.Player) (float64, int) {
	maximizingPlayer := false
	if currentPlayer == MAXIMIZINGPLAYER {
		maximizingPlayer = true
	}
	return _minmax(position, depth, maximizingPlayer, -1)
}

type Game struct {
	Board       [][]string
	PlayersTurn misc.Player
	Width       int
	Height      int
}

// PrintHelp diplays a helper message
func (g *Game) PrintHelp() {
	fmt.Printf("\n%v\n", strings.Repeat("-", len(g.Board[0])*4+1))
	out := "|"
	for i := range g.Board[0] {
		out += misc.NoPlayerValue
		out += strconv.Itoa(i)
		out += " |"
	}
	fmt.Println(out)
}

// copyBoard returns a new game with a board identical to g
func (g *Game) copyBoard() Game {
	newGame := GetNewGame()
	copiedBoard := make([][]string, len(g.Board))
	for i := range g.Board {
		copiedBoard[i] = make([]string, len(g.Board[i]))
		copy(copiedBoard[i], g.Board[i])
	}
	newGame.Board = copiedBoard
	return newGame
}

// AlternatePlayersTurn alternates the player whose move it is
func (g *Game) AlternatePlayersTurn() {
	if g.PlayersTurn == 0 {
		g.PlayersTurn = 1
		return
	}
	g.PlayersTurn = 0
}

// PrintBoard pretty prints the board
func (g *Game) PrintBoard() {
	fmt.Println(strings.Repeat("-", len(g.Board[0])*4+1))
	for i, _ := range g.Board {
		out := "|"
		for j, _ := range g.Board[i] {
			out += misc.NoPlayerValue
			out += g.Board[i][j]
			out += " |"
		}
		fmt.Println(out)
		fmt.Println(strings.Repeat("-", len(g.Board[0])*4+1))
	}
}

// init initializes the board and sets all fields to misc.NoPlayerValue
func (g *Game) init() {
	g.PlayersTurn = 1
	g.Width = 6
	g.Height = 8
	g.Board = make([][]string, g.Height)
	for i, _ := range g.Board {
		g.Board[i] = make([]string, g.Width)
		for j, _ := range g.Board[i] {
			g.Board[i][j] = misc.NoPlayerValue
		}
	}
}

func (g *Game) _selectComputerMove() int {
	_, column := minmax(*g, 7, g.PlayersTurn)
	return column
}

// selectComputerMove is a wraper for _selectComputerMove
func (g *Game) selectComputerMove() int {
	fmt.Println("Computer move")
	time.Sleep(0 * time.Second)
	return g._selectComputerMove()
}

// selectHumanMove gets a move from the user
func (g *Game) selectHumanMove(reader *bufio.Reader) (int, error) {
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

// SelectMove wraps doComputerMove and doHumanMove
func (g *Game) SelectMove(reader *bufio.Reader) int {
	if g.PlayersTurn == 0 {
		num, err := g.selectHumanMove(reader)
		if err != nil {
			fmt.Println("Invalid input, please try again")
			g.SelectMove(reader)
		}
		return num
	}
	return g.selectComputerMove()
}

// fall simulates the falling of the token in the column
// It returns the row in which the token will rest or
// an error if this column is already full
func (g *Game) fall(column int) (int, error) {
	if g.Board[0][column] != misc.NoPlayerValue {
		return 0, fmt.Errorf("column %v is already full", column)
	}
	for row := 0; row < len(g.Board); row++ {
		if g.Board[row][column] != misc.NoPlayerValue {
			return row - 1, nil
		}
	}
	return len(g.Board) - 1, nil
}

// DoMove does executes a selected move on the board
// returns an error if the selected column is already full
func (g *Game) DoMove(column int) error {
	row, err := g.fall(column)
	if err != nil {
		// fmt.Printf("Column %v already full. Choose again\n", column)
		// TODO: Activate, except during minmax
		return err
	}
	g.Board[row][column] = misc.PlayerToString(g.PlayersTurn)
	return nil
}

// checkKHorizontal checks if the next k horizontal to the right
// fields are equal to playerString
// Returns false, if out of bound
// If xPos is 0, the fields 0, 1, 2, and 3 will be looked at
func (g *Game) checkKHorizontal(playerSting string, k int, xPos int, yPos int) bool {
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
func (g *Game) checkKVertical(playerSting string, k int, xPos int, yPos int) bool {
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

// checkKDiagonalDown checks if the next k diagonal to the bottom right
// fields are equal to playerString
// Returns false, if out of bound
// If yPos and xPos are both 0, the fields 1,1, 2,2, and 3,3 will be looked at
func (g *Game) checkKDiagonalDown(playerSting string, k int, xPos int, yPos int) bool {
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
func (g *Game) checkKDiagonalUp(playerSting string, k int, xPos int, yPos int) bool {
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
func (g *Game) checkX(playerString string, k int, xPos int, yPos int, d misc.Direction) bool {
	if d == misc.Horizontal {
		return g.checkKHorizontal(playerString, k, xPos, yPos)
	} else if d == misc.Vertical {
		return g.checkKVertical(playerString, k, xPos, yPos)
	} else if d == misc.Diagonal {
		return g.checkKDiagonalDown(playerString, k, xPos, yPos) ||
			g.checkKDiagonalUp(playerString, k, xPos, yPos)
	}
	panic(fmt.Sprintf("Unsupported direction %v", d))
}

// Won returns the winning player and true
// if no player Won, it returns false
func (g *Game) Won() (bool, misc.Player) {
	// iterate over each pos in board
	for yPos := 0; yPos < g.Height; yPos++ {
		for xPos := 0; xPos < g.Width; xPos++ {
			// call each direction
			for dir := 0; dir < 3; dir++ {
				//call each player
				for i := 0; i <= 1; i++ {
					if g.Board[yPos][xPos] != misc.NoPlayerValue &&
						g.checkX(misc.PlayerIntToStrig(i), 4, xPos, yPos, misc.Direction(dir)) {
						return true, misc.Player(i)
					}
				}
			}
		}
	}

	return false, -1
}

// BoardFull returns true if no more moves are possible
func (g *Game) BoardFull() bool {
	for _, val := range g.Board[0] {
		if val == misc.NoPlayerValue {
			return false
		}
	}
	return true
}

func GetNewGame() Game {
	var g Game
	g.init()
	return g
}
