package main

import (
	"fmt"
)

type Board struct {
	rowOne   []string
	rowTwo   []string
	rowThree []string
}

func (b *Board) returnDiagonals() ([]string, []string) {
	diagonalOne := []string{
		b.rowOne[0],
		b.rowTwo[1],
		b.rowThree[2],
	}
	diagonalTwo := []string{
		b.rowOne[2],
		b.rowTwo[1],
		b.rowThree[0],
	}

	return diagonalOne, diagonalTwo
}

func (b *Board) returnColumns() ([]string, []string, []string) {
	columnOne := []string{
		b.rowOne[0],
		b.rowTwo[0],
		b.rowThree[0],
	}
	columnTwo := []string{
		b.rowOne[1],
		b.rowTwo[1],
		b.rowThree[1],
	}
	columnThree := []string{
		b.rowOne[2],
		b.rowTwo[2],
		b.rowThree[2],
	}

	return columnOne, columnTwo, columnThree
}

func (b *Board) updateBoard(x int, y int, player string) {
	if y == 0 {
		b.rowOne[x] = player
	} else if y == 1 {
		b.rowTwo[x] = player
	} else {
		b.rowThree[x] = player
	}
}

func (b *Board) legalMove(x int, y int) bool {
	if y > 2 || x > 2 {
		fmt.Println("Please ensure x & y values are <= 2")
		return false
	}

	if y == 0 {
		if b.rowOne[x] != " " {
			fmt.Printf("Piece %s already at location", b.rowOne[x])
			return false
		}
	}

	if y == 1 {
		if b.rowTwo[x] != " " {
			fmt.Printf("Piece %s already at location", b.rowTwo[x])
			return false
		}
	}

	if y == 2 {
		if b.rowThree[x] != " " {
			fmt.Printf("Piece %s already at location", b.rowThree[x])
			return false
		}
	}

	return true
}

func (b *Board) displayBoard() {
	printArray(b.rowOne[:])
	fmt.Printf("\n-----------\n")
	printArray(b.rowTwo[:])
	fmt.Println("\n-----------")
	printArray(b.rowThree[:])
	fmt.Println("")
}

func printArray(x []string) {
	for i := 0; i < len(x); i++ {
		if i == 1 {
			fmt.Printf("| %s |", x[i])
		} else {
			fmt.Printf(" %s ", x[i])
		}
	}
}
