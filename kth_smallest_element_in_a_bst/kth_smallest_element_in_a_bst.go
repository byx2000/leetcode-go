package kth_smallest_element_in_a_bst

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/

import . "leetcode-go/common"

func kthSmallest(root *TreeNode, k int) (ret int) {
	defer func() {
		n := recover()
		if n != nil {
			ret = n.(int)
		}
	}()

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)

		if k == 1 {
			panic(node.Val)
		}
		k--

		traverse(node.Right)
	}

	traverse(root)
	return 0
}
