package main

import (
	"fmt"
)

func controller(b *Board) {
	gameState := false
	player := "X"

	for !gameState {
		var x, y int

		fmt.Println("Enter x,y coordinates with a space between them (Ex: 1 1): ")
		fmt.Scanln(&x, &y)

		if b.legalMove(x, y) {
			b.updateBoard(x, y, player)

			if checkWin(b, player) {
				fmt.Printf("%s is the winner!\n", player)
				gameState = true
				b.displayBoard()
				return
			}

			b.displayBoard()

			player = playerOpposite(player)
		}
	}
}
