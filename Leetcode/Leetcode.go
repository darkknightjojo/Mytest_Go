package Leetcode

import ()

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	preHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	var prev *ListNode
	prev = preHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}

	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}

	return preHead.Next
}
