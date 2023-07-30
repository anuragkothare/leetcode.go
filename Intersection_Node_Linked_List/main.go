// https://leetcode.com/problems/intersection-of-two-linked-lists/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Create the linked list one: 1 -> 2 -> 4 -> 5
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 4}
	node4 := &ListNode{Val: 5}

	head.Next = node2
	node2.Next = node3
	node3.Next = node4

	// Create the linked list two: 1 -> 3 -> 4 -> 5 (Intersection with list 1 at common node 4)
	head2 := &ListNode{Val: 1}
	node6 := &ListNode{Val: 3}

	head2.Next = node6
	node6.Next = node3

	printLinkedList(head)
	printLinkedList(head2)

	interSectionNode := getIntersectionNode(head, head2)

	// Print the interSection Node
	fmt.Println(interSectionNode)
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	d1 := headA
	d2 := headB
	listLengthDiff := 0

	headA_Length := listLength(headA)
	headB_Length := listLength(headB)

	if headA_Length > headB_Length {
		d1 = headA
		d2 = headB
		listLengthDiff = headA_Length - headB_Length
	} else {
		d1 = headB
		d2 = headA
		listLengthDiff = headB_Length - headA_Length
	}

	for listLengthDiff != 0 {
		d1 = d1.Next
		listLengthDiff--
	}

	for d1 != nil && d2 != nil {
		if d1 == d2 {
			return d1
		}
		d1 = d1.Next
		d2 = d2.Next
	}
	return nil
}

// Optimal Approch
func getIntersectionNode2(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}
	d1 := headA
	d2 := headB

	for d1 != d2 {
		if d1 == nil {
			d1 = headB
		} else {
			d1 = d1.Next
		}

		if d2 == nil {
			d2 = headA
		} else {
			d2 = d2.Next
		}
	}
	return d1
}

func listLength(l1 *ListNode) int {
	count := 0
	current := l1

	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}
