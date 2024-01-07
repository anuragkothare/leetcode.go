package main

import "fmt"

func main() {
	grid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	m := len(grid)
	n := len(grid[0])
	ans := uniquePathsWithObstaclesSP(m, n, grid)
	fmt.Println(ans)
}

func uniquePathsWithObstacles(row int, col int, grid [][]int) int {
	dp := make([][]int, row)

	for i := range dp {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				dp[0][0] = 1
			} else if grid[i][j] == 1 {
				dp[i][j] = 0
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

// Space optimization solution
func uniquePathsWithObstaclesSP(row int, col int, grid [][]int) int {
	prev := make([]int, col)

	for i := 0; i < row; i++ {
		curr := make([]int, col)
		for j := 0; j < col; j++ {
			if i >= 0 && j >= 0 && grid[i][j] == 1 {
				curr[j] = 0
			} else if i == 0 && j == 0 {
				curr[0] = 1
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
