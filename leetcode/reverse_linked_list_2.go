package leetcode

/*
92. Reverse Linked List II (Medium)
Given the head of a singly linked list and two integers left and right where left <= right,
reverse the nodes of the list from position left to position right, and return the reversed list.


Example 1:
Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]

Example 2:
Input: head = [5], left = 1, right = 1
Output: [5]


Constraints:
- The number of nodes in the list is n.
- 1 <= n <= 500
- -500 <= Node.val <= 500
- 1 <= left <= right <= n

Follow up: Could you do it in one pass?
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	var (
		cur  = head
		prev *ListNode
		next *ListNode
	)

	var preReverse *ListNode
	i := 1
	for ; i < left; i++ {
		if i+1 == left {
			preReverse = cur
		}

		cur = cur.Next
	}

	reverseStart := cur
	for ; i <= right; i++ {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	reverseEnd := prev

	if preReverse == nil && cur == nil {
		return reverseEnd
	}
	if preReverse == nil {
		reverseStart.Next = cur
		return reverseEnd
	}

	preReverse.Next = reverseEnd
	reverseStart.Next = cur

	return head
}
