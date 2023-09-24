// https://leetcode.com/problems/spiral-matrix/description/

package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	ans := spiralOrder(arr)
	fmt.Println(ans)

}

func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	left := 0
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1

	for left <= right && top <= bottom {
		// right traverse
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++

		// bottom traverse
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--

		// left traverse
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			bottom--
		}

		// up
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}

	return ans

}
