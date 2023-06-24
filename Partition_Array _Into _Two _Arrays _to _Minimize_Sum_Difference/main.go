// https://leetcode.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/description/

// Solution works for non-negative elements in array only.

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{3, 9, 7, 3}
	res := minimumDifference(nums)
	fmt.Println(res)
}

func minimumDifference(nums []int) int {
	arr_length := len(nums)
	arr_sum := 0

	for _, val := range nums {
		arr_sum += val
	}
	fmt.Println(arr_sum, arr_length)

	min_diff := math.MaxInt
	for s1 := 0; s1 <= arr_sum/2; s1++ {
		isTarget := subsetSumToK(arr_length, s1, &nums)
		if isTarget {
			s2 := arr_sum - s1
			min_diff = min(min_diff, int(math.Abs(float64(s1-s2))))
		}
	}
	return min_diff
}

func subsetSumToK(n int, k int, arr *[]int) bool {
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, k+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if (*arr)[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-(*arr)[i-1]]
			}
		}
	}
	return dp[n][k]
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
