package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1}
	ans := findMaxConsecutiveOnes(nums)
	fmt.Println(ans)
}

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count = 0
		} else {
			count++
			maxCount = maxInt(maxCount, count)
		}
	}
	return maxCount
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
