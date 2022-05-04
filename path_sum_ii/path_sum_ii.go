package path_sum_ii

// https://leetcode-cn.com/problems/path-sum-ii/

import (
	. "leetcode-go/common"
)

func pathSum(root *TreeNode, targetSum int) [][]int {
	var paths [][]int
	var path []int

	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if node.Val == sum {
				path = append(path, node.Val)
				p := make([]int, len(path))
				copy(p, path)
				paths = append(paths, p)
				path = path[:len(path)-1]
			}
			return
		}
		path = append(path, node.Val)
		dfs(node.Left, sum-node.Val)
		dfs(node.Right, sum-node.Val)
		path = path[:len(path)-1]
	}

	dfs(root, targetSum)
	return paths
}
