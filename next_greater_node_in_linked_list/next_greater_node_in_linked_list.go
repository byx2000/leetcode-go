package next_greater_node_in_linked_list

// https://leetcode.cn/problems/next-greater-node-in-linked-list/

import . "leetcode-go/common"

func nextLargerNodes(head *ListNode) []int {
	// 将链表转换成切片
	toList := func(head *ListNode) []int {
		var list []int
		for head != nil {
			list = append(list, head.Val)
			head = head.Next
		}
		return list
	}

	nums := toList(head)
	var result = make([]int, len(nums))

	// dp[i]表示nums[i]右边比nums[i]大的第一个元素的索引
	var dp = make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		idx := i + 1
		for idx != 0 && nums[idx] <= nums[i] {
			idx = dp[idx]
		}
		dp[i] = idx
		if dp[i] == 0 {
			result[i] = 0
		} else {
			result[i] = nums[dp[i]]
		}
	}

	return result
}
