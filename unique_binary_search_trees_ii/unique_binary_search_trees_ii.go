package unique_binary_search_trees_ii

// https://leetcode.cn/problems/unique-binary-search-trees-ii/

import . "leetcode-go/common"

func generateTrees(n int) []*TreeNode {
	// 生成节点值为i~j的所有二叉搜索树
	var build func(i, j int) []*TreeNode
	build = func(i, j int) []*TreeNode {
		if i > j {
			return []*TreeNode{nil}
		}

		if i == j {
			return []*TreeNode{{i, nil, nil}}
		}

		var nodes []*TreeNode

		// k为根节点，[i, k-1]为左子树，[k+1, j]为右子树
		for k := i; k <= j; k++ {
			leftNodes := build(i, k-1)
			rightNodes := build(k+1, j)
			for _, left := range leftNodes {
				for _, right := range rightNodes {
					nodes = append(nodes, &TreeNode{k, left, right})
				}
			}
		}

		return nodes
	}

	return build(1, n)
}
