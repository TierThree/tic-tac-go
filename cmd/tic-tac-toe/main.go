package main

func main() {
	board := &Board{
		rowOne:   []string{" ", " ", " "},
		rowTwo:   []string{" ", " ", " "},
		rowThree: []string{" ", " ", " "},
	}

	board.displayBoard()

	controller(board)
}
