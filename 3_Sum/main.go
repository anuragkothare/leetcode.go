// https://leetcode.com/problems/3sum

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			curr_sum := nums[i] + nums[j] + nums[k]
			if curr_sum < 0 {
				j++
			} else if curr_sum > 0 {
				k--
			} else {
				temp := []int{nums[i], nums[j], nums[k]}
				ans = append(ans, temp)
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return ans

}
