package validate_binary_search_tree

// https://leetcode-cn.com/problems/validate-binary-search-tree/

import (
	. "leetcode-go/common"
	"math"
)

func isValidBST1(root *TreeNode) (ret bool) {
	var last int64 = math.MinInt64

	defer func() {
		e := recover()
		if e != nil {
			ret = false
		}
	}()

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)

		if int64(node.Val) <= last {
			panic(0)
		}
		last = int64(node.Val)

		traverse(node.Right)
	}

	traverse(root)

	return true
}

func isValidBST2(root *TreeNode) bool {
	var last int64 = math.MinInt64

	var traverse func(node *TreeNode) bool
	traverse = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !traverse(node.Left) {
			return false
		}
		if int64(node.Val) <= last {
			return false
		}
		last = int64(node.Val)
		return traverse(node.Right)
	}

	return traverse(root)
}
