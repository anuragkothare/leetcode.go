package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	max, sub_arr := maxSubArray(nums)
	fmt.Println("Maximum sub array sum is:", max)
	fmt.Println("Maximum sum sub array is:", sub_arr)

}

// Maximum Sub array sum using Kadane Algorithm
func maxSubArray(nums []int) (int, []int) {
	start_ind := -1
	end_ind := -1
	sum := 0
	maxi := math.MinInt

	for i := 0; i < len(nums); i++ {
		if sum == 0 {
			start_ind = i
		}
		sum += nums[i]
		if sum > maxi {
			maxi = sum
			end_ind = i
		}
		if sum < 0 {
			sum = 0
		}
	}
	sub_arr := nums[start_ind : end_ind+1]
	return maxi, sub_arr
}
