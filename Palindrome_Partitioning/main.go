// https://leetcode.com/problems/palindrome-partitioning/
package main

import "fmt"

func main() {
	s := "cbbbcc"
	ans := partition(s)
	fmt.Println(ans)
}

func partition(str string) [][]string {
	res := [][]string{}
	path := []string{}

	helper(str, 0, path, &res)
	return res
}

func helper(s string, ind int, path []string, res *[][]string) {
	if ind == len(s) {
		*res = append(*res, append([]string{}, path...))
		return
	}

	for i := ind; i < len(s); i++ {
		if isPalindrome(s[ind : i+1]) {
			path = append(path, s[ind:i+1])
			helper(s, i+1, path, res)
			path = path[:len(path)-1]
		}
	}
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
