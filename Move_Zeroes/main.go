package main

import "fmt"

func main() {
	input := []int{1, 0, 0, 5, 9, 0, 0, 7, 0, 0, 8}
	moveZeroes(input)
	fmt.Println(input)
}

func moveZeroes(nums []int) {
	// Initialize a pointer (j) to keep track of the position to swap non-zero elements
	j := 0

	// Iterate through the array
	for i := 0; i < len(nums); i++ {
		// If the current element is non-zero
		if nums[i] != 0 {
			// Swap the non-zero element with the element at position j
			nums[i], nums[j] = nums[j], nums[i]

			// Increment the pointer to the next position for potential non-zero elements
			j++
		}
	}
}
