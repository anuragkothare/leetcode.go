// https: //leetcode.com/problems/remove-duplicates-from-sorted-list/description/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Create the linked list: 1 -> 1 -> 3 -> 3 -> 5
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 5}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	printLinkedList(head)

	// Remove Duplicated from the Linked List
	head = deleteDuplicates(head)

	// Print the updated linked list
	printLinkedList(head)
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
			if curr.Next != nil && curr.Val != curr.Next.Val {
				curr = curr.Next
			}
		} else {
			curr = curr.Next
		}
	}
	return head
}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}
