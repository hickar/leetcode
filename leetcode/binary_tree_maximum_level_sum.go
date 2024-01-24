package leetcode

import "fmt"

func maxLevelSum(root *TreeNode) int {
	var levels []int
	breadthFirstTraversal(root, func(item TreeNodeItem) {
		if item.Depth > len(levels)-1 {
			levels = append(levels, item.Node.Val)
			return
		}

		levels[item.Depth] += item.Node.Val
	})

	return findMaxIdx(levels) + 1
}

func findMaxIdx(nums []int) int {
	maxIdx := 0

	for i, item := range nums {
		fmt.Println(i, item)
		if item > nums[maxIdx] {
			maxIdx = i
		}
	}

	return maxIdx
}
