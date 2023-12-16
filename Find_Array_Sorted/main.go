package main

import "fmt"

func main() {
	nums := []int{1, 4, 5, 6, 9, 9, 17}

	res := isArraySorted(nums)
	fmt.Println(res)
}

func isArraySorted(nums []int) bool {
	isSorted := true
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			isSorted = false
			return isSorted
		}
	}
	return isSorted
}
