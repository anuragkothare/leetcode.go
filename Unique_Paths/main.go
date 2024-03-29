// https://leetcode.com/problems/unique-paths/

package main

import "fmt"

func main() {
	row := 3
	col := 3
	res := uniquePaths2(row, col)
	fmt.Println(res)
}

// Recursion + Memoization
func uniquePathsRec(row, col int) int {
	dp := make([][]int, row)

	for i := range dp {
		dp[i] = make([]int, col)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return helper(row-1, col-1, &dp)
}

func helper(i, j int, dp *[][]int) int {
	if i == 0 && j == 0 {
		return 1
	}
	if (*dp)[i][j] != -1 {
		return (*dp)[i][j]
	}
	up, down := 0, 0
	if i > 0 {
		up = helper(i-1, j, dp)
	}
	if j > 0 {
		down = helper(i, j-1, dp)
	}
	(*dp)[i][j] = up + down
	return (*dp)[i][j]
}

// Tabulation Method
func uniquePaths(row int, col int) int {
	dp := make([][]int, row)

	for i := range dp {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				dp[0][0] = 1
			} else {
				up, left := 0, 0
				if i > 0 {
					up = dp[i-1][j]
				}
				if j > 0 {
					left = dp[i][j-1]
				}
				dp[i][j] = up + left
			}
		}
	}
	return dp[row-1][col-1]
}

// DP with Space Optimization
func uniquePaths2(row int, col int) int {
	prev := make([]int, col)

	for i := 0; i < row; i++ {
		curr := make([]int, col)
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				curr[j] = 1
			} else {
				up, left := 0, 0
				if i > 0 {
					up = prev[j]
				}
				if j > 0 {
					left = curr[j-1]
				}
				curr[j] = up + left
			}
		}
		prev = curr
	}
	return prev[col-1]
}
