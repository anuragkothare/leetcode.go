package main

// https://leetcode.com/problems/permutations/

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	ans := permute(nums)
	fmt.Println(ans)
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}

	var out [][]int

	for i, val := range nums {
		var copyNum = make([]int, len(nums)-1)
		copy(copyNum, nums[:i])
		copy(copyNum[i:], nums[i+1:])

		for _, line := range permute(copyNum) {
			out = append(out, append([]int{val}, line...))
		}
	}

	return out
}
