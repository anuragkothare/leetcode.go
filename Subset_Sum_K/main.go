// https://www.codingninjas.com/codestudio/problems/subset-sum-equal-to-k_1550954

package main

import "fmt"

func main() {
	arr := []int{1, 7, 2, 9, 10}
	n := len(arr)
	k := 6
	res := subsetSumToK(n, k, &arr)
	fmt.Println("Subset sum to", k, "exists:", res)
}

// Tabulation
func subsetSumToK(n int, k int, arr *[]int) bool {
	// dp[i][j] will represent whether it is possible to achieve the sum j using the first i elements of the nums array.
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, k+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if (*arr)[i-1] > k {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-(*arr)[i-1]]
			}
		}
	}
	return dp[n][k]
}
