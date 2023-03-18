// https://leetcode.com/problems/subsets-ii/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 2}
	ans := subsetsWithDup(arr)
	fmt.Println(ans)
}

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	helper(nums, 0, []int{}, &res)
	return res
}

func helper(nums []int, ind int, subset []int, res *[][]int) {
	temp := make([]int, len(subset))
	copy(temp, subset)
	*res = append(*res, temp)
	for i := ind; i < len(nums); i++ {
		if i > 0 && i > ind && nums[i] == nums[i-1] {
			continue
		}
		subset = append(subset, nums[i])
		helper(nums, i+1, subset, res)
		subset = subset[:len(subset)-1]
	}
}
