package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func branchSums(root *TreeNode) []int {
	var sums []int

	var calculateSums func(node *TreeNode, currentSum int)
	calculateSums = func(node *TreeNode, currentSum int) {
		if node == nil {
			return
		}

		currentSum += node.Value

		// If it's a leaf node, append the current sum to the result slice
		if node.Left == nil && node.Right == nil {
			sums = append(sums, currentSum)
			return
		}

		// Recursively calculate sums for left and right subtrees
		calculateSums(node.Left, currentSum)
		calculateSums(node.Right, currentSum)
	}

	calculateSums(root, 0)
	return sums
}

func main() {
	// Create a binary tree
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}
	root.Left.Left = &TreeNode{Value: 4}
	root.Left.Right = &TreeNode{Value: 5}

	// Calculate branch sums
	sums := branchSums(root)
	fmt.Println(sums) // Output: [7 8 4]
}
