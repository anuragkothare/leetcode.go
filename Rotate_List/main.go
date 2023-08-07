// https://leetcode.com/problems/rotate-list/

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

	k := 7

	printLinkedList(head)

	// Rotate Right the linked list
	rotateRightK := rotateRight(head, k)

	printLinkedList(rotateRightK)
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// compute the length of the list
	cur := head
	list_len := 1
	for cur.Next != nil {
		list_len++
		cur = cur.Next
	}
	cur.Next = head
	k = k % list_len
	k = list_len - k

	for k > 0 {
		k--
		cur = cur.Next

	}

	head = cur.Next
	cur.Next = nil

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
