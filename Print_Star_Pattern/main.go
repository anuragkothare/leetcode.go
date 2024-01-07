package main

import "fmt"

func main() {
	printStarPattern(15)
}

func printStarPattern(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()

	}
}
