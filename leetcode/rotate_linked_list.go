package leetcode

/*
61. Rotate List (Medium)
Given the head of a linked list, rotate the list to the right by k places.


Example 1:
Input: head = [1,2,3,4,5], k = 2
Output: [4,5,1,2,3]

Example 2:
Input: head = [0,1,2], k = 4
Output: [2,0,1]


Constraints:
- The number of nodes in the list is in the range [0, 500].
- -100 <= Node.val <= 100
- 0 <= k <= 2 * 109
*/

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	var (
		count       int
		cur         = head
		rotateStart *ListNode
		rotateEnd   *ListNode
	)

	for cur != nil {
		count++
		rotateEnd = cur
		cur = cur.Next
	}
	if k%count == 0 {
		return head
	}

	cur = head
	for i := 0; i < count-(k%count)-1; i++ {
		cur = cur.Next
	}
	rotateStart = cur.Next
	cur.Next = nil

	rotateEnd.Next = head
	return rotateStart
}
