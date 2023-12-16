package main

import "fmt"

func main() {
	nums := []int{2, 4, 5, 8, 4, 7}
	res := secondLargestElement(nums)
	fmt.Println(res)
}

func secondLargestElement(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	firstLargest := nums[0]
	secondLargest := nums[1]

	if secondLargest > firstLargest {
		firstLargest, secondLargest = secondLargest, firstLargest
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > firstLargest {
			secondLargest = firstLargest
			firstLargest = nums[i]
		} else if nums[i] > secondLargest && nums[i] != firstLargest {
			secondLargest = nums[i]
		}
	}
	return secondLargest
}
