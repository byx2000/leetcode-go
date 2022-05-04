package lowest_common_ancestor_of_a_binary_tree

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

import . "leetcode-go/common"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 判断节点n是否在二叉树node内
	var in func(node *TreeNode, n *TreeNode) bool
	in = func(node *TreeNode, n *TreeNode) bool {
		if node == nil {
			return false
		}
		if node == n {
			return true
		}
		return in(node.Left, n) || in(node.Right, n)
	}

	if root == p || root == q {
		return root
	}

	b1 := in(root.Left, p)
	b2 := in(root.Left, q)
	b3 := in(root.Right, p)
	b4 := in(root.Right, q)

	if b1 && b2 {
		return lowestCommonAncestor(root.Left, p, q)
	} else if b3 && b4 {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}
