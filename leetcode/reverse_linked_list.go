package leetcode

/*
206. Reverse Linked List (Easy)
Given the head of a singly linked list, reverse the list,
and return the reversed list.


Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
  if head == nil || head.Next == nil {
    return head
  }

  var (
    cur = head
    prev *ListNode
    next *ListNode
  )

  for cur != nil {
    next = cur.Next
    cur.Next = prev
    prev = cur

    cur = next
  }

  return prev
}
