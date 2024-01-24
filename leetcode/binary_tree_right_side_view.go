package leetcode

import "github.com/hickar/leetcode/structures"

/*
199. Binary Tree Right Side View (Medium)
Given the root of a binary tree, imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.


Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

Example 2:
Input: root = [1,null,3]
Output: [1,3]

Example 3:
Input: root = []
Output: []
*/

type TreeNodeItem struct {
	Depth int
	Node  *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var levels [][]int

	breadthFirstTraversal(root, func(item TreeNodeItem) {
		if item.Depth > len(levels)-1 {
			levels = append(levels, []int{item.Node.Val})
			return
		}

		levels[item.Depth] = append(levels[item.Depth], item.Node.Val)
	})

	rightSideView := make([]int, 0, len(levels))
	for _, level := range levels {
		rightSideView = append(rightSideView, level[len(level)-1])
	}

	return rightSideView
}

func breadthFirstTraversal(root *TreeNode, callback func(TreeNodeItem)) {
	queue := structures.NewQueue[TreeNodeItem](10)
	queue.Enqueue(TreeNodeItem{
		Node:  root,
		Depth: 0,
	})

	for queue.Len() > 0 {
		item, ok := queue.Dequeue()
		if !ok || item.Node == nil {
			break
		}
		callback(item)

		if item.Node.Left != nil {
			queue.Enqueue(TreeNodeItem{
				Node:  item.Node.Left,
				Depth: item.Depth + 1,
			})
		}
		if item.Node.Right != nil {
			queue.Enqueue(TreeNodeItem{
				Node:  item.Node.Right,
				Depth: item.Depth + 1,
			})
		}
	}
}
