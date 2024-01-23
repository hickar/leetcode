package leetcode

/*
  328. Odd Even Linked List (Medium)
  Given the head of a singly linked list, group all the nodes with odd indices together
  followed by the nodes with even indices, and return the reordered list.
  The first node is considered odd, and the second node is even, and so on.
  Note that the relative order inside both the even and odd groups should remain as it was in the input.
  You must solve the problem in O(1) extra space complexity and O(n) time complexity.


  Example 1:
  Input: head = [1,2,3,4,5]
  Output: [1,3,5,2,4]

  Example 2:
  Input: head = [2,1,3,5,6,4,7]
  Output: [2,3,6,7,1,5,4]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	var (
		cur      = head
		next     *ListNode
		headOdd  *ListNode
		headEven *ListNode
		curOdd   *ListNode
		curEven  *ListNode
	)

	for i := 1; ; i++ {
		if cur == nil {
			break
		}
		next = cur.Next
		cur.Next = nil

		switch {
		case i%2 == 0 && curEven != nil:
			curEven.Next = cur
			curEven = curEven.Next

		case i%2 == 0 && curEven == nil:
			headEven = cur
			curEven = headEven

		case i%2 == 1 && curOdd != nil:
			curOdd.Next = cur
			curOdd = curOdd.Next

		case i%2 == 1 && curOdd == nil:
			headOdd = cur
			curOdd = headOdd
		}

		cur = next
	}

	if curOdd != nil {
		curOdd.Next = headEven
	}

	return head
}
