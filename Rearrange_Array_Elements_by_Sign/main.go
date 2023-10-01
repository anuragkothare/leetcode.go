// https://leetcode.com/problems/rearrange-array-elements-by-sign

package main

import "fmt"

func main() {
	nums := []int{3, 1, -2, -5, 2, -4, -1, -7}
	res := rearrangeArray3(nums)
	fmt.Println("res", res)
}

// Brute Force Approach
func rearrangeArray(nums []int) []int {
	pos_arr := make([]int, 0)
	neg_arr := make([]int, 0)
	res_arr := make([]int, 0)

	for i := range nums {
		if nums[i] >= 0 {
			pos_arr = append(pos_arr, nums[i])
		} else {
			neg_arr = append(neg_arr, nums[i])
		}
	}

	p_ind := 0
	n_ind := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 && len(pos_arr) > 0 {
			res_arr = append(res_arr, pos_arr[p_ind])
			p_ind++
		} else if i%2 != 0 && len(neg_arr) > 0 {
			res_arr = append(res_arr, neg_arr[n_ind])
			n_ind++
		}
	}
	return res_arr
}

// Optimal with one iteration/pass
func rearrangeArray2(nums []int) []int {
	res := make([]int, len(nums))
	pos_ind := 0
	neg_ind := 1

	for i := range nums {
		if nums[i] >= 0 {
			res[pos_ind] = nums[i]
			pos_ind += 2
		} else {
			res[neg_ind] = nums[i]
			neg_ind += 2
		}
	}
	return res
}

// 2nd Variety if number of positive elements is not equal to negative elements
func rearrangeArray3(nums []int) []int {
	pos_arr := make([]int, 0)
	neg_arr := make([]int, 0)

	for i := range nums {
		if nums[i] >= 0 {
			pos_arr = append(pos_arr, nums[i])
		} else {
			neg_arr = append(neg_arr, nums[i])
		}
	}

	if len(pos_arr) > len(neg_arr) {
		for i := 0; i < len(neg_arr); i++ {
			nums[2*i] = pos_arr[i]
			nums[2*i+1] = neg_arr[i]
		}

		var index int = len(neg_arr) - 2
		for i := len(neg_arr); i < len(pos_arr); i++ {
			nums[index] = pos_arr[i]
			index++
		}
	} else {
		for i := 0; i < len(pos_arr); i++ {
			nums[2*i] = pos_arr[i]
			nums[2*i+1] = neg_arr[i]
		}

		var index int = len(pos_arr) - 2
		for i := len(pos_arr); i < len(neg_arr); i++ {
			nums[index] = neg_arr[i]
			index++
		}
	}
	return nums
}
