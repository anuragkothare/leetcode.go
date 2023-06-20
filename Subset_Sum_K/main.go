// https://www.codingninjas.com/codestudio/problems/subset-sum-equal-to-k_1550954

package main

import "fmt"

func main() {
	arr := []int{1, 7, 2, 9, 10}
	n := len(arr)
	k := 6
	res := subsetSumToK2(n, k, &arr)
	fmt.Println("Subset sum to", k, "exists:", res)
}

// Tabulation
func subsetSumToK(n int, k int, arr *[]int) bool {
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, k+1)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = true
		dp[0][(*arr)[0]] = true

		for ind := 1; ind < n; ind++ {
			for target := 1; target <= k; target++ {
				notTake := dp[ind-1][target]
				take := false

				if (*arr)[ind] <= target {
					take = dp[ind-1][target-(*arr)[ind]]
				}
				dp[ind][target] = take || notTake

			}
		}
	}
	return dp[n-1][k]
}

// Space Optimization of Tabulation
func subsetSumToK2(n int, k int, arr *[]int) bool {
	fmt.Println(*arr)
	prev := make([]bool, k+1)
	curr := make([]bool, k+1)

	prev[0] = true
	curr[0] = true
	prev[(*arr)[0]] = true

	for ind := 1; ind < n; ind++ {
		for target := 1; target <= k; target++ {
			notTake := prev[target]
			take := false

			if (*arr)[ind] <= target {
				take = prev[target-(*arr)[ind]]
			}
			curr[target] = take || notTake
		}
		prev = curr
	}

	return prev[k]
}
