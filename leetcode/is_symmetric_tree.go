package leetcode

import "github.com/hickar/leetcode/structures"

/*
101. Symmetric Tree (Easy)
Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

Example 1:
Input: root = [1,2,2,3,4,4,3]
Output: true

Example 2:
Input: root = [1,2,2,null,3,null,3]
Output: false

*/

func IsSymmetric(root *structures.TreeNode[int]) bool {
	if root == nil {
		return true
	}

	return areMirrored(root.Left, root.Right)
}

func areMirrored(l, r *structures.TreeNode[int]) bool {
	if l == nil && r == nil {
		return true
	}
	if (l != nil && r == nil) || (l == nil && r != nil) {
		return false
	}
	if l.Value != r.Value {
		return false
	}

	return areMirrored(l.Left, r.Right) && areMirrored(l.Right, r.Left)
}
