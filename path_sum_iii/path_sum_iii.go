package path_sum_iii

// https://leetcode-cn.com/problems/path-sum-iii/

import . "leetcode-go/common"

func pathSum(root *TreeNode, targetSum int) int {
	var count func(node *TreeNode, sum int) int
	count = func(node *TreeNode, sum int) int {
		if node == nil {
			return 0
		}

		if node.Val == sum {
			return 1 + count(node.Left, sum-node.Val) + count(node.Right, sum-node.Val)
		}

		return count(node.Left, sum-node.Val) + count(node.Right, sum-node.Val)
	}

	if root == nil {
		return 0
	}

	return count(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}
