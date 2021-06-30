package main

import (
	"bufio"
	"fmt"
	"fourwins/main/src/game"
	"fourwins/main/src/misc"
	"os"
)

func main() {
	fmt.Printf("Let's play 4 wins!\nThe top line indicates the rows you can choose\n")
	g := game.GetNewGame()

	ioReader := bufio.NewReader(os.Stdin)

	// game loop
	for !g.BoardFull() {
		// players do action
		columnSelected := g.SelectMove(ioReader)
		err := g.DoMove(columnSelected)
		if err != nil {
			continue
		}
		g.PrintHelp()
		g.PrintBoard()
		gameFinished, winner := g.Won()
		if gameFinished {
			fmt.Printf("Congrats to player %v. You are a winner!\n", misc.PlayerToString(winner))
			return
		}
		g.AlternatePlayersTurn()
	}
	fmt.Println("Board full. Game ends")
}
