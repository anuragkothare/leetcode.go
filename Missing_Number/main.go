package main

import "fmt"

func main() {
	input_arr := []int{3, 0, 1}

	res := missingNumberOptimalXOR(input_arr)
	fmt.Println(res)
}

func missingNumber(nums []int) int {
	n := len(nums)
	res := make([]bool, n+1)

	for i := 0; i < n; i++ {
		res[nums[i]] = true
	}

	for j := 0; j < len(res); j++ {
		if res[j] == false {
			return j
		}
	}
	return -1
}

// Optimal solution by eleminating extra array of slices of boolean
func missingNumberOptimal(nums []int) int {
	n := len(nums)
	expectedSum := n * (n + 1) / 2
	actualSum := 0

	for i := 0; i < n; i++ {
		actualSum += nums[i]
	}
	ans := expectedSum - actualSum
	return ans
}

// Optimal solution by XOR
func missingNumberOptimalXOR(nums []int) int {
	n := len(nums)

	expectedXOR := 0

	// for i := 0; i <= n; i++ {
	// 	expectedXOR ^= i
	// }

	actualXOR := 0

	for i := 0; i < n; i++ {
		actualXOR ^= nums[i]
		expectedXOR ^= i
	}
	expectedXOR = expectedXOR ^ n
	return expectedXOR ^ actualXOR
}
