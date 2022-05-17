package jump_game

// https://leetcode.cn/problems/jump-game/

func canJump(nums []int) bool {
	// 当前位于nums[i]，还能走step步
	step := 0
	for i := 0; i < len(nums); i++ {
		if step < nums[i] {
			step = nums[i]
		}

		if step == 0 && i != len(nums)-1 {
			return false
		}

		step--
	}

	return true
}
