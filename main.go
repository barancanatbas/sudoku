package main

import (
	"fmt"
)

var rows, cols, boxes [9][10]bool

func findEmptyCells(board [][]int) [][2]int {
	var emptyCells [][2]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				emptyCells = append(emptyCells, [2]int{i, j})
			} else {
				num := board[i][j]
				rows[i][num] = true
				cols[j][num] = true
				boxes[boxIndex(i, j)][num] = true
			}
		}
	}
	return emptyCells
}

func boxIndex(row, col int) int {
	return (row/3)*3 + col/3
}

func getPossibleValues(row, col int) []int {
	var possible []int
	for num := 1; num <= 9; num++ {
		if !rows[row][num] && !cols[col][num] && !boxes[boxIndex(row, col)][num] {
			possible = append(possible, num)
		}
	}
	return possible
}

func updatePossibleValues(board [][]int, possibleValues map[[2]int][]int) {
	for pos := range possibleValues {
		row, col := pos[0], pos[1]
		possibleValues[pos] = getPossibleValues(row, col)
	}
}

func setSingleValues(board [][]int, possibleValues map[[2]int][]int) bool {
	changed := false
	for pos, values := range possibleValues {
		if len(values) == 1 {
			num := values[0]
			row, col := pos[0], pos[1]
			board[row][col] = num
			rows[row][num] = true
			cols[col][num] = true
			boxes[boxIndex(row, col)][num] = true
			delete(possibleValues, pos)
			changed = true
		}
	}
	return changed
}

func solveSudoku(board [][]int) bool {
	emptyCells := findEmptyCells(board)
	possibleValues := make(map[[2]int][]int)
	for _, cell := range emptyCells {
		row, col := cell[0], cell[1]
		possibleValues[cell] = getPossibleValues(row, col)
	}

	for {
		updatePossibleValues(board, possibleValues)
		if !setSingleValues(board, possibleValues) {
			break
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func printBoard(board [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	board := [][]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	if solveSudoku(board) {
		printBoard(board)
	} else {
		fmt.Println("Error: Sudoku could not be solved.")
	}
}
