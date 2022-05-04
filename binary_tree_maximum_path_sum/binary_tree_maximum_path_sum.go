package binary_tree_maximum_path_sum

// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/

import (
	. "leetcode-go/common"
	"math"
)

func maxPathSum(root *TreeNode) int {
	max := func(n int, nums ...int) int {
		r := n
		for _, num := range nums {
			r = int(math.Max(float64(r), float64(num)))
		}
		return r
	}

	var maxSum func(node *TreeNode) int
	maxSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(node.Val, node.Val+maxSum(node.Left), node.Val+maxSum(node.Right))
	}

	r := root.Val

	r1 := maxSum(root.Left)
	if r1 > 0 {
		r += r1
	}

	r2 := maxSum(root.Right)
	if r2 > 0 {
		r += r2
	}

	if root.Left != nil {
		r = max(r, maxPathSum(root.Left))
	}
	if root.Right != nil {
		r = max(r, maxPathSum(root.Right))
	}
	return r
}
