package leetcode

func deepestLeavesSumOptimzed(root *TreeNode) int {
	sum := 0
	maxDepth := 0

	var dfs func(n *TreeNode, depth int)

	dfs = func(n *TreeNode, depth int) {
		if n == nil {
			return
		}

		if depth > maxDepth {
			sum = 0
			maxDepth = depth
		}

		if depth == maxDepth {
			sum += n.Val
		}

		dfs(n.Left, depth+1)
		dfs(n.Right, depth+1)
	}

	dfs(root, 0)

	return sum
}

func deepestLeavesSum(root *TreeNode) int {
	var deepestLeaves []TreeNodeItem

	breadthFirstTraversal(root, func(item TreeNodeItem) {
		if len(deepestLeaves) > 0 && deepestLeaves[0].Depth < item.Depth {
			deepestLeaves = nil
		}

		deepestLeaves = append(deepestLeaves, item)
	})

	sum := 0
	for i := 0; i < len(deepestLeaves); i++ {
		sum += deepestLeaves[i].Node.Val
	}

	return sum
}
