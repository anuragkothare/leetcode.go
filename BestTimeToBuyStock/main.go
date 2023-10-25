// https://leetcode.com/problems/best-time-to-buy-and-sell-stock

package main

import "fmt"

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	max_profit := maxProfit(nums)
	fmt.Println(max_profit)
}

func maxProfit(prices []int) int {
	minBuyPrice := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		cost := prices[i] - minBuyPrice
		profit = maxElement(profit, cost)
		minBuyPrice = minElement(minBuyPrice, prices[i])
	}
	return profit
}

func maxElement(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minElement(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
