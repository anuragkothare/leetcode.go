// https://leetcode.com/problems/climbing-stairs/

package main

import "fmt"

func main() {
	steps := 3
	res := climbStairs(steps)
	fmt.Println(res)
}

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
