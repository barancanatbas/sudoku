package main

import "fmt"

var rows, cols, boxes [9][10]bool

func solveSudoku(board [][]int) bool {
	emptyCells := findEmptyCells(board)
	return backtrack(board, emptyCells, 0)
}

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

func backtrack(board [][]int, emptyCells [][2]int, index int) bool {
	if index == len(emptyCells) {
		return true
	}

	row, col := emptyCells[index][0], emptyCells[index][1]

	for num := 1; num <= 9; num++ {
		if isValid(row, col, num) {
			board[row][col] = num
			rows[row][num] = true
			cols[col][num] = true
			boxes[boxIndex(row, col)][num] = true

			if backtrack(board, emptyCells, index+1) {
				return true
			}

			board[row][col] = 0
			rows[row][num] = false
			cols[col][num] = false
			boxes[boxIndex(row, col)][num] = false
		}
	}

	return false
}

func isValid(row, col, num int) bool {
	return !rows[row][num] && !cols[col][num] && !boxes[boxIndex(row, col)][num]
}

func boxIndex(row, col int) int {
	return (row/3)*3 + col/3
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
		fmt.Println("there is no solution")
	}
}
