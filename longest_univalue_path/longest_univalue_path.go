package longest_univalue_path

// https://leetcode-cn.com/problems/longest-univalue-path/

import (
	. "leetcode-go/common"
	"math"
)

func longestUnivaluePath(root *TreeNode) int {
	max := func(n int, nums ...int) int {
		r := n
		for _, num := range nums {
			r = int(math.Max(float64(r), float64(num)))
		}
		return r
	}

	var count func(node *TreeNode, val int) int
	count = func(node *TreeNode, val int) int {
		if node == nil {
			return 0
		}
		if node.Val == val {
			return 1 + max(count(node.Left, val), count(node.Right, val))
		}
		return 0
	}

	if root == nil {
		return 0
	}

	len1 := count(root.Left, root.Val)
	len2 := count(root.Right, root.Val)
	return max(len1+len2, longestUnivaluePath(root.Left), longestUnivaluePath(root.Right))
}
