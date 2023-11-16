package leetcode

/*
141. Linked List Cycle (Easy) (???)

Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be reached again
by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to.
Note that pos is not passed as a parameter.
Return true if there is a cycle in the linked list. Otherwise, return false.
*/

func HasCycle(head *ListNode) bool {
	memo := make(map[int]map[*ListNode]struct{})
	cur := head

	for cur != nil {
		_, ok := memo[cur.Val]
		if ok {
			if _, ok = memo[cur.Val][cur]; ok {
				return true
			}
		} else {
			memo[cur.Val] = make(map[*ListNode]struct{}, 1)
		}

		memo[cur.Val][cur] = struct{}{}
		cur = cur.Next
	}

	return false
}
