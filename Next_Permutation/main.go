package main

import "fmt"

func main() {
	nums := []int{1, 3, 2}
	res := nextPermutation(nums)
	fmt.Println(res)

}

func nextPermutation(nums []int) []int {
	ind := -1
	nums_length := len(nums)

	for i := nums_length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			ind = i
			break
		}
	}
	if ind == -1 {
		return reverseSlice(nums)
	}
	for i := nums_length - 1; i > ind; i-- {
		if nums[i] > nums[ind] {
			swap(&nums[i], &nums[ind])
			break
		}
	}
	subArr := reverseSlice(nums[ind+1:])
	nums = nums[:ind+1]
	return append(nums, subArr...)
}

func reverseSlice(arr []int) []int {
	i := 0
	j := len(arr) - 1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
