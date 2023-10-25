package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{2, 3, -2, 4}
	res := maxProduct(nums)
	fmt.Println(res)
}

// On^2 Solution
func maxProduct(nums []int) int {
	n := len(nums)
	maxProduct := math.MinInt
	for i := 0; i < n; i++ {
		product := 1
		for j := i; j < n; j++ {
			product = product * nums[j]
			maxProduct = maxInt(product, maxProduct)
		}

	}
	return maxProduct
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
