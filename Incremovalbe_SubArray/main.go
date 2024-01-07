package main

import "fmt"

func main() {
	arr := []int{8, 7, 6, 6}

	res := incremovableSubarrayCount(arr)
	fmt.Println(res)
}

func incremovableSubarrayCount(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			last := -1
			flag := 1
			for k := 0; k < len(nums); k++ {
				if k >= i && k <= j {
					continue
				}
				if last >= nums[k] {
					flag = 0
					break
				}
				last = nums[k]
			}
			if flag == 1 {
				ans++
			}
		}
	}
	return ans
}
