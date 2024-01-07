// https://leetcode.com/discuss/interview-question/4291609/Amazon-or-SDE-Online-Assessment/

package main

import (
	"fmt"
)

func getMaximumIds(inp []int) []int {
	result := make([]int, 0)
	maxVal := 0
	seenMap := make(map[int]int)

	for _, num := range inp {
		if num > 0 {
			seenMap[num]++
			if seenMap[num] > maxVal {
				maxVal = seenMap[num]
			}
		} else {
			if seenMap[-num] == maxVal {
				maxVal--
			}
			seenMap[-num]--
		}
		result = append(result, maxVal)
	}

	return result
}

func main() {
	arr := []int{1, 2, -1, 2} // Not working Case
	result := getMaximumIds(arr)
	fmt.Println(result)
}
