// https://leetcode.com/problems/fibonacci-number

package main

import "fmt"

func main() {
	n := 7
	res := fib3(n)
	fmt.Println(res)
}

// Using Memoization
// Time complexity: O(n)
// Space complexity: O(n) + O(n)

func fib1(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	if dp[n] != -1 {
		return dp[n]
	}
	dp[n] = fib1(n-1) + fib1(n-2)
	return dp[n]
}

// Using Tabulation
// Time complexity: O(n)
// Space complexity: O(n)
func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// Iterative Solution / Space Optimized Soln.
// Time complexity: O(n)
// Space complexity: O(1)
func fib3(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	prev2, prev := 0, 1

	for i := 2; i <= n; i++ {
		curr := prev2 + prev
		prev2 = prev
		prev = curr
	}
	return prev
}
