package main

import "fmt"

func main() {
	nums := []int{4, 2, 1, 1, 3}
	res := findMajorityElement(nums)
	fmt.Println(res)
}

// You can find the majority element in an array with constant space complexity O(1) and linear time complexity O(n) using the Boyer-Moore Majority Vote algorithm. This algorithm is efficient and straightforward to implement. It works by maintaining a candidate element and a count. Here's a Go implementation of this algorithm:

func findMajorityElement(nums []int) int {
	var element int
	var count int

	for _, val := range nums {
		if count == 0 {
			element = val
			count++
		} else if element == val {
			count++
		} else {
			count--
		}
	}

	// At this point, the 'candidate' should be the majority element.
	// You can perform a second pass to verify if it appears more than n/2 times,
	// but this step is not required as the problem statement assumes there is
	// always a majority element.

	return element
}
