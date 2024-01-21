package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	ans := [][]int{}

	// Sort intervals based on their start times
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < n; i++ {
		start, end := intervals[i][0], intervals[i][1]

		if len(ans) > 0 && end <= ans[len(ans)-1][1] {
			continue
		}

		for j := i + 1; j < n; j++ {
			if intervals[j][0] <= end {
				end = max(end, intervals[j][1])
			} else {
				break
			}
		}

		ans = append(ans, []int{start, end})
	}

	return ans
}

// Optimized
func mergeOptimized(intervals [][]int) [][]int {
	n := len(intervals)
	ans := [][]int{}

	// Sort intervals based on their start times
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < n; i++ {
		if len(ans) == 0 || (intervals[i][0] > ans[len(ans)-1][1]) {
			ans = append(ans, intervals[i])
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], intervals[i][1])
		}
	}
	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result := mergeOptimized(intervals)
	fmt.Println(result)
}
