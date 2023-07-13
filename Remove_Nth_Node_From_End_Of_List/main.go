// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

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

	n := 2

	printLinkedList(head)

	// Remove the Nth node from the end
	head = removeNthFromEnd(head, n)

	// Print the updated linked list
	printLinkedList(head)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Create a dummy node to handle the case of removing the head
	dummy := &ListNode{Val: 0, Next: head}

	fast := dummy
	slow := dummy

	// Move the fast pointer N steps forward
	for i := 1; i <= n; i++ {
		fast = fast.Next
	}

	// Move both pointers simultaneously

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Remove the Nth node by updating the Next pointer of slow
	slow.Next = slow.Next.Next

	return dummy.Next

}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()

}
