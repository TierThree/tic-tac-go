package main

func main() {
	board := &Board{
		rowOne:   [3]string{" ", " ", " "},
		rowTwo:   [3]string{" ", " ", " "},
		rowThree: [3]string{" ", " ", " "},
	}

	board.displayBoard()

	controller(board)
}
