// https://leetcode.com/problems/triangle/description/

package main

import (
	"fmt"
)

func main() {
	// Example usage:
	triangleArray := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	result := minimumTotalSP(triangleArray)
	fmt.Println(result)
}

// Recursion Memoization
func minimumTotal(triangle [][]int) int {
	tri_length := len(triangle)
	dp := make([][]int, tri_length)

	row := 0
	for row < tri_length {
		col := make([]int, row+1)
		for c := range col {
			col[c] = -1
		}
		dp[row] = col
		row++
	}
	return helper(0, 0, triangle, tri_length, &dp)
}

func helper(i int, j int, tri [][]int, n int, dp *[][]int) int {
	if i == n-1 {
		return tri[i][j]
	}

	if (*dp)[i][j] != -1 {
		return (*dp)[i][j]
	}

	down := tri[i][j] + helper(i+1, j, tri, n, dp)
	diagonal := tri[i][j] + helper(i+1, j+1, tri, n, dp)

	(*dp)[i][j] = min(down, diagonal)
	return (*dp)[i][j]
}

// Tabulation Approach
func minimumTotalSP(triangle [][]int) int {
	m := len(triangle)

	dp := make([][]int, m)

	row := 0

	for row <= m-1 {
		col := make([]int, row+1)
		dp[row] = col
		row++
	}

	for c := m - 1; c >= 0; c-- {
		dp[m-1][c] = triangle[m-1][c]
	}

	for i := m - 2; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			dp[i][j] = min(triangle[i][j]+dp[i+1][j+1], triangle[i][j]+dp[i+1][j])
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
