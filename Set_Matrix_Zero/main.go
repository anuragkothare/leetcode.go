package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	rowLength := len(matrix)
	colLength := len(matrix[0])

	zeroRowMap, zeroColMap := make(map[int]bool), make(map[int]bool)

	// Identify the positions of zeroes
	for r := 0; r < rowLength; r++ {
		for c := 0; c < colLength; c++ {
			if matrix[r][c] == 0 {
				zeroRowMap[r] = true
				zeroColMap[c] = true
			}
		}
	}

	// Update the matrix
	for row := range zeroRowMap {
		for col := 0; col < colLength; col++ {
			matrix[row][col] = 0
		}
	}

	for col := range zeroColMap {
		for row := 0; row < rowLength; row++ {
			matrix[row][col] = 0
		}
	}
}

// Space Optimized Solution
func setZeroesSpaceOptimized(matrix [][]int) {
	rowLength := len(matrix)
	colLength := len(matrix[0])

	colZero := 1

	// Identify the positions of zeroes
	for r := 0; r < rowLength; r++ {
		for c := 0; c < colLength; c++ {
			if matrix[r][c] == 0 {
				matrix[r][0] = 0
				if c == 0 {
					colZero = 0
				} else {
					matrix[0][c] = 0
				}
			}
		}
	}

	for r := rowLength - 1; r >= 1; r-- {
		for c := colLength - 1; c >= 1; c-- {
			if matrix[r][c] != 0 {
				if matrix[0][c] == 0 || matrix[r][0] == 0 {
					matrix[r][c] = 0
				}
			}
		}
	}

	if matrix[0][0] == 0 {
		for col := 0; col < colLength; col++ {
			matrix[0][col] = 0
		}
	}

	if colZero == 0 {
		for row := 0; row < rowLength; row++ {
			matrix[row][0] = 0
		}
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}

	fmt.Println("Original Matrix:")
	printMatrix(matrix)

	setZeroesSpaceOptimized(matrix)

	fmt.Println("\nMatrix after setting zeroes:")
	printMatrix(matrix)
}
