// https://leetcode.com/problems/linked-list-cycle/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Create the linked list: 1 -> 2 -> 3 -> 4 -> 5 ---> 2 Cycle
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node2

	isCycle := hasCycle(head)

	// Is Cycle present
	if isCycle {
		fmt.Println("Cycle Detected")
	} else {
		fmt.Println("Cycle not Detected")
	}
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	walker := head
	runner := head

	for runner != nil && runner.Next != nil {
		walker = walker.Next
		runner = runner.Next.Next

		if walker == runner {
			return true
		}
	}
	return false
}
