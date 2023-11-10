package main

import "fmt"

func main() {
	newGame := true
	continuation := "Y"

	for newGame {
		board := &Board{
			rowOne:   []string{" ", " ", " "},
			rowTwo:   []string{" ", " ", " "},
			rowThree: []string{" ", " ", " "},
		}

		fmt.Println("New game [Yn]?")
		fmt.Scanln(&continuation)

		if continuation == "Y" {
			board.displayBoard()
			controller(board)
		} else {
			fmt.Println("Have a good time!")
		}
	}
}
