// https://www.codingninjas.com/studio/problems/ninja-s-training_3621003
package main

import "fmt"

func main() {
	// Example usage
	pointsArray := [][]int{{1, 2, 5}, {3, 1, 1}, {3, 3, 3}}
	result := maxMeritPointsSP(pointsArray)
	fmt.Println(result)
}

// Tabulation (Bottom_Up)
func maxMeritPoints(points [][]int) int {
	n := len(points)

	dp := make([][]int, n)
	for i := range dp {
		row := make([]int, 4)
		dp[i] = row
	}

	dp[0][0] = max(points[0][1], points[0][2])
	dp[0][1] = max(points[0][0], points[0][1])
	dp[0][2] = max(points[0][0], points[0][1])
	dp[0][3] = max(points[0][0], max(points[0][1], points[0][2]))

	for day := 1; day < n; day++ {
		for last := 0; last < 4; last++ {
			for task := 0; task < 3; task++ {
				if task != last {
					point := points[day][task] + dp[day-1][task]
					dp[day][last] = max(dp[day][last], point)
				}
			}
		}
	}
	return dp[n-1][3]
}

// Space Optimized Solution
func maxMeritPointsSP(points [][]int) int {
	n := len(points)

	dp := make([]int, 4) // Create a single 1-D array
	dp[0] = max(points[0][1], points[0][2])
	dp[1] = max(points[0][0], points[0][1])
	dp[2] = max(points[0][0], points[0][1])
	dp[3] = max(points[0][0], max(points[0][1], points[0][2]))

	for day := 1; day < n; day++ {
		temp := make([]int, 4)
		for last := 0; last < 4; last++ {
			for task := 0; task < 3; task++ {
				if task != last {
					point := points[day][task] + dp[task]
					temp[last] = max(temp[last], point)
				}
			}
		}
		dp = temp
	}
	return dp[3]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
