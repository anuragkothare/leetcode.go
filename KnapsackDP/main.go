package main

import "fmt"

func main() {
	values := []int{60, 100, 120}
	weights := []int{10, 20, 30}
	capacity := 50

	maxValue, includedItems := knapsackDP(values, weights, capacity)
	fmt.Println("Maximum value:", maxValue)
	fmt.Println("Included items:", includedItems)
}

func knapsackDP(values []int, weights []int, capacity int) (int, []int) {
	n := len(values)
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}
	// dp[i][w] represents the maximum value that can be achieved using the first i items (from item 0 to item i-1) and a knapsack capacity of w.
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if w >= weights[i-1] {
				dp[i][w] = max(values[i-1]+dp[i-1][w-weights[i-1]], dp[i-1][w])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// Traceback to find the items included in the knapsack
	includedItems := make([]int, 0)
	i, w := n, capacity

	for i > 0 && w > 0 {
		if dp[i][w] != dp[i-1][w] {
			includedItems = append(includedItems, i-1)
			w -= weights[i-1]
		}
		i--
	}
	return dp[n][capacity], reverse(includedItems)

}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
