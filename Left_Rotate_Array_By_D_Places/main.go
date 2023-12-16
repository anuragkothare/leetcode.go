package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	res := leftRotateArrayOptimal(nums, 3)
	fmt.Println(res)
}

func leftRotateArray(nums []int, d int) []int {
	if d <= 0 || d >= len(nums) {
		// No rotation needed for invalid or full rotations
		return nums
	}

	temp := make([]int, d)
	copy(temp, nums[:d])

	for i := d; i < len(nums); i++ {
		nums[i-d] = nums[i]
	}

	copy(nums[len(nums)-d:], temp)

	return nums
}

// Optimal Solution
func leftRotateArrayOptimal(nums []int, d int) []int {
	if d <= 0 || d >= len(nums) {
		// No rotation needed for invalid or full rotations
		return nums
	}

	reverseSlice(&nums, 0, d-1)
	reverseSlice(&nums, d, len(nums)-1)
	reverseSlice(&nums, 0, len(nums)-1)
	return nums
}

func reverseSlice(nums *[]int, start int, end int) {
	for start <= end {
		temp := (*nums)[start]
		(*nums)[start] = (*nums)[end]
		(*nums)[end] = temp
		start++
		end--
	}
}
