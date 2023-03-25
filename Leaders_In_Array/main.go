// https://www.codingninjas.com/codestudio/problems/superior-elements_6783446?utm_source=youtube&utm_medium=affiliate&utm_campaign=striver_Arrayproblems

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	inputArr := []int{1, 2, 3, 2}
	res := superiorElements(inputArr)
	fmt.Println(res)
}

func superiorElements(input []int) []int {
	leadersList := []int{}
	min := math.MinInt32
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] > min {
			leadersList = append(leadersList, input[i])
			min = input[i]
		}
	}
	sort.Ints(leadersList)
	return leadersList
}
