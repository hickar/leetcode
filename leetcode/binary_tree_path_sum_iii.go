package leetcode

/*
437. Path Sum III (Medium)
Given the root of a binary tree and an integer targetSum,
return the number of paths where the sum of the values along the path equals targetSum.
The path does not need to start or end at the root or a leaf,
but it must go downwards (i.e., traveling only from parent nodes to child nodes).


Example 1:
Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
Output: 3
Explanation: The paths that sum to 8 are shown.

Example 2:
Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: 3
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TODO: Think about better solution.

func pathSum(root *TreeNode, targetSum int) int {
	var pathCount int

	inorderTraversal(root, func(node *TreeNode) {
		pathCount += countSumPaths(node, 0, targetSum)
	})

	return pathCount
}

func countSumPaths(root *TreeNode, curSum, targetSum int) int {
	if root == nil {
		return 0
	}

	pathCount := 0
	curSum += root.Val
	if curSum == targetSum {
		pathCount += 1
	}

	return pathCount + countSumPaths(root.Left, curSum, targetSum) + countSumPaths(root.Right, curSum, targetSum)
}

func inorderTraversal(root *TreeNode, callback func(*TreeNode)) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, callback)
	callback(root)
	inorderTraversal(root.Right, callback)
}
