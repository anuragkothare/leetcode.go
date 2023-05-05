// https://leetcode.com/problems/house-robber/https://leetcode.com/problems/house-robber/

package main

import "fmt"

func main() {
	input_arr := []int{2, 7, 9, 3, 1}
	res := rob1(input_arr)
	fmt.Println(res)
}

// DP Bottom Up Approach Tabulation
func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		pick := nums[i]
		if i > 1 {
			pick += dp[i-2]
		}
		no_pick := 0 + dp[i-1]
		dp[i] = maxInt(pick, no_pick)
	}
	return dp[n-1]
}

// DP Bottom Up Approach Tabulation Space Optimization
func rob1(nums []int) int {
	n := len(nums)
	prev2, prev := 0, nums[0]

	for i := 1; i < n; i++ {
		pick := nums[i]
		if i > 1 {
			pick += prev2
		}
		no_pick := 0 + prev
		cur := maxInt(pick, no_pick)
		prev2 = prev
		prev = cur
	}
	return prev
}

func maxInt(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
