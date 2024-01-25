package leetcode

/*
102. Binary Tree Level Order Traversal (Medium)
Given the root of a binary tree, return the level order traversal
of its nodes' values. (i.e., from left to right, level by level).


Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []
*/

func levelOrder(root *TreeNode) [][]int {
	var levels [][]int

	breadthFirstTraversal(root, func(item TreeNodeItem) {
		if item.Depth > len(levels)-1 {
			levels = append(levels, []int{item.Node.Val})
			return
		}

		levels[item.Depth] = append(levels[item.Depth], item.Node.Val)
	})

	return levels
}
