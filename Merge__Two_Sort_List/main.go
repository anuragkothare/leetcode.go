// https://leetcode.com/problems/merge-two-sorted-lists/

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

	// Create the linked list two: 1 -> 3 -> 4
	head2 := &ListNode{Val: 1}
	node6 := &ListNode{Val: 3}
	node7 := &ListNode{Val: 4}

	head2.Next = node6
	node6.Next = node7

	printLinkedList(head)
	printLinkedList(head2)

	// Merge two Sortd Linked Lists
	mergedList := mergeTwoLists(head, head2)

	// Print the merged linked list
	printLinkedList(mergedList)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list2
	}

	if list1.Val > list2.Val {
		temp := list1
		list1 = list2
		list2 = temp
	}
	res := list1

	for list1 != nil && list2 != nil {
		tmp := &ListNode{}
		for list1 != nil && list1.Val <= list2.Val {
			tmp = list1
			list1 = list1.Next
		}
		tmp.Next = list2

		temp := list1
		list1 = list2
		list2 = temp

	}
	return res

}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}
	fmt.Println()
}
