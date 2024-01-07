// https://leetcode.com/problems/longest-consecutive-sequence

package main

import "fmt"

func main() {
	nums := []int{100, 102, 105, 103, 101, 104, 4, 200, 1, 3, 2}
	result := longestConsecutive(nums)
	fmt.Println("Longest Consecutive Sequence Length:", result)
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 1

	for num := range numSet {
		// Check if num is the start of a sequence
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			// Expand the sequence
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			// Update the maximum length
			if currentStreak > maxLength {
				maxLength = currentStreak
			}
		}
	}

	return maxLength
}
