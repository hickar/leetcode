package leetcode

import "github.com/hickar/leetcode/structures"

/*
100. Same Tree (Easy)

Given the roots of two binary trees p and q, write a function to check if they are the same or not.
Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.

Example 1:
Input: p = [1,2,3], q = [1,2,3]
Output: true

Example 2:
Input: p = [1,2], q = [1,null,2]
Output: false

Example 3:
Input: p = [1,2,1], q = [1,1,2]
Output: false

*/

func IsSameTree(p *structures.TreeNode[int], q *structures.TreeNode[int]) bool {
	if (p != nil && q == nil) || (p == nil && q != nil) {
		return false
	}
	if p == nil && q == nil {
		return true
	}

	return p.Value == q.Value &&
		IsSameTree(p.Left, q.Left) &&
		IsSameTree(p.Right, q.Right)
}
