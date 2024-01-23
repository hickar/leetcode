package leetcode

/*
872. Leaf-Similar Trees (Easy)
Consider all the leaves of a binary tree, from left to right order,
the values of those leaves form a leaf value sequence.
For example, in the given tree above,
the leaf value sequence is (6, 7, 4, 9, 8).
Two binary trees are considered leaf-similar
if their leaf value sequence is the same.
Return true if and only if the two given trees
with head nodes root1 and root2 are leaf-similar.


Example 1:
Input:
  root1 = [3,5,1,6,2,9,8,null,null,7,4],
  root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
Output: true

Example 2:
Input:
  root1 = [1,2,3],
  root2 = [1,3,2]
Output: false
/*

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
  var (
    leaves1 = gatherLeafValues(root1)
    leaves2 = gatherLeafValues(root2)
  )

  if len(leaves1) != len(leaves2) {
    return false
  }

  for i := 0; i < len(leaves1); i++ {
    if leaves1[i] != leaves2[i] {
      return false
    }
  }

  return true
}

func gatherLeafValues(root *TreeNode) []int {
  var leafArr []int

  if root == nil {
    return nil
  }

  if root.Left == nil && root.Right == nil {
    return []int{root.Val}
  }

  leafArr = append(leafArr, gatherLeafValues(root.Left)...)
  leafArr = append(leafArr, gatherLeafValues(root.Right)...)

  return leafArr
}
