package main

import (
	"golang.org/x/exp/slices"
)

func checkWin(b *Board, player string) bool {
	pOpp := playerOpposite(player)

	diagOne, diagTwo := b.returnDiagonals()
	colOne, colTwo, colThree := b.returnColumns()

	lines := []([]string){
		b.rowOne,
		b.rowTwo,
		b.rowThree,
		diagOne,
		diagTwo,
		colOne,
		colTwo,
		colThree,
	}

	for i := 0; i < len(lines); i++ {
		if !(slices.Contains(lines[i], pOpp)) && !(slices.Contains(lines[i], " ")) {
			return true
		}
	}

	return false
}

func playerOpposite(player string) string {
	if player == "X" {
		return "O"
	}

	return "X"
}

func fullBoard(b *Board) bool {
	if !(slices.Contains(b.rowOne, " ")) &&
		!(slices.Contains(b.rowTwo, " ")) &&
		!(slices.Contains(b.rowThree, " ")) {
		return true
	}

	return false
}
