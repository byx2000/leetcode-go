package house_robber_iii

// https://leetcode.cn/problems/house-robber-iii/

import (
	. "leetcode-go/common"
	"math"
)

func rob1(root *TreeNode) int {
	var dp func(root *TreeNode, takeRoot bool) int
	dp = func(root *TreeNode, takeRoot bool) int {
		if root == nil {
			return 0
		}

		if takeRoot {
			return root.Val + dp(root.Left, false) + dp(root.Right, false)
		} else {
			return int(math.Max(float64(dp(root.Left, true)), float64(dp(root.Left, false)))) +
				int(math.Max(float64(dp(root.Right, true)), float64(dp(root.Right, false))))
		}
	}

	return int(math.Max(float64(dp(root, true)), float64(dp(root, false))))
}

func rob2(root *TreeNode) int {
	cache := map[*TreeNode]int{}

	var dp func(root *TreeNode) int
	dp = func(root *TreeNode) (ret int) {
		if v, exist := cache[root]; exist {
			return v
		}

		defer func() {
			cache[root] = ret
		}()

		if root == nil {
			return 0
		}

		// 不偷根节点
		r1 := dp(root.Left) + dp(root.Right)

		// 偷根节点
		r2 := root.Val
		if root.Left != nil {
			r2 += dp(root.Left.Left) + dp(root.Left.Right)
		}
		if root.Right != nil {
			r2 += dp(root.Right.Left) + dp(root.Right.Right)
		}

		return int(math.Max(float64(r1), float64(r2)))
	}

	return dp(root)
}
