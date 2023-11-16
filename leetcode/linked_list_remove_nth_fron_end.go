package leetcode

/*
19. Remove Nth Node From End of List (Medium)
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]
*/

// CPU optimized version

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := make([]*ListNode, 0)

	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}

	idx := len(nodes) - n
	left, right := idx-1, idx+1

	switch {
	case left > -1 && right < len(nodes):
		nodes[left].Next = nodes[right]
	case left > -1 && right > len(nodes)-1:
		nodes[left].Next = nil
	case left < 0 && right < len(nodes):
		head = nodes[right]
	default:
		return nil
	}

	return head
}
