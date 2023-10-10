package leetcode

import (
	"fmt"
	"math"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbersLinkedList(l1 *ListNode, l2 *ListNode) *ListNode {
	node := &ListNode{0, nil}
	nodePtr := node
	remainder := 0

	for {
		if l1 != nil && l2 != nil {
			nodePtr.Val = l1.Val + l2.Val + remainder
			remainder = 0
		} else if l1 != nil {
			nodePtr.Val = l1.Val + remainder
			remainder = 0
		} else if l2 != nil {
			nodePtr.Val = l2.Val + remainder
			remainder = 0
		} else {
			nodePtr.Val = remainder
		}

		if nodePtr.Val >= 10 {
			nodePtr.Val -= 10
			remainder = 1
		} else if nodePtr.Val == 9 && remainder == 1 {
			remainder = 1
		} else if nodePtr.Val == 9 && remainder == 0 {
			remainder = 0
		} else {
			remainder = 0
		}

		if l1 != nil && l2 != nil {
			l1, l2 = l1.Next, l2.Next
		} else if l1 != nil {
			l1, l2 = l1.Next, nil
		} else if l2 != nil {
			l1, l2 = nil, l2.Next
		}

		if l1 == nil && l2 == nil {
			if remainder == 1 {
				nodePtr.Next = &ListNode{remainder, nil}
			} else {
				nodePtr.Next = nil
			}
			break
		} else {
			nodePtr.Next = &ListNode{}
		}

		nodePtr = nodePtr.Next
	}

	return node
}

func PrintNodeList(node *ListNode) {
	result := 0
	for i := 0; node != nil; i++ {
		result += int(math.Pow(float64(10), float64(i))) * node.Val
		node = node.Next
	}
	fmt.Println(result)
}
