package Leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	node14 := ListNode{
		Val:  4,
		Next: nil,
	}
	node12 := ListNode{
		Val:  2,
		Next: &node14,
	}
	node11 := ListNode{
		Val:  1,
		Next: &node12,
	}
	node23 := ListNode{
		Val:  3,
		Next: nil,
	}
	node22 := ListNode{
		Val:  2,
		Next: &node23,
	}
	node21 := ListNode{
		Val:  1,
		Next: &node22,
	}
	fmt.Println(mergeTwoLists(&node11, &node21))
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
