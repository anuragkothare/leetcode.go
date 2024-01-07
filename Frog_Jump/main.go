// https://www.codingninjas.com/studio/problems/frog-jump_3621012

package main

import (
	"fmt"
	"math"
)

func main() {
	N := 4
	heights := []int{10, 20, 30, 10}
	result := minEnergyToReachN(N, heights)
	fmt.Println("Minimum energy needed:", result)
}

func minEnergyToReachN(N int, heights []int) int {
	if N == 1 {
		return 0
	}
	dp := make([]int, N)

	for i := range dp {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0

	for i := 1; i < N; i++ {
		dp[i] = min(dp[i], dp[i-1]+abs(heights[i-1]-heights[i]))
		if i > 1 {
			dp[i] = min(dp[i], dp[i-2]+abs(heights[i-2]-heights[i]))
		}
	}
	return dp[N-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
