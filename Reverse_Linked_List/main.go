// https://leetcode.com/problems/reverse-linked-list/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Create the linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	printLinkedList(head)

	// Reverse the linked list
	reversedHead := reverseList(head)

	printLinkedList(reversedHead)
}

func reverseList(head *ListNode) *ListNode {
	var curr, prev, next *ListNode
	curr = head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}
