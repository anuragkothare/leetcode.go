package main

import "fmt"

func main() {
	nums := []int{1, 3, 3, 5, 5, 8, 8, 9, 9, 10}
	res := removeDuplicates(nums)
	fmt.Println(res)
}

func removeDuplicates(nums []int) []int {
	cur := 0
	next := 1

	for next < len(nums) {
		if nums[cur] != nums[next] {
			cur++
			nums[cur] = nums[next]
		}
		next++
	}
	return nums[:cur+1]
}
