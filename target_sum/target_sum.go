package target_sum

// https://leetcode-cn.com/problems/target-sum/

func findTargetSumWays1(nums []int, target int) (cnt int) {
	var dfs func(index int, sum int)
	dfs = func(index int, sum int) {
		if index == len(nums) {
			if sum == target {
				cnt++
			}
			return
		}

		dfs(index+1, sum+nums[index])
		dfs(index+1, sum-nums[index])
	}

	dfs(0, 0)
	return
}

func findTargetSumWays2(nums []int, target int) int {
	// dp(index, sum) 表示nums[0]~nums[index]和为sum共有多少种方法
	var dp func(index, sum int) int
	dp = func(index, sum int) int {
		if index == 0 {
			cnt := 0
			if nums[0] == sum {
				cnt++
			}
			if -nums[0] == sum {
				cnt++
			}
			return cnt
		}
		return dp(index-1, sum+nums[index]) + dp(index-1, sum-nums[index])
	}

	return dp(len(nums)-1, target)
}

func findTargetSumWays3(nums []int, target int) int {
	cache := make([][]int, len(nums))
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, 4001)
		for j := 0; j < len(cache[i]); j++ {
			cache[i][j] = -1
		}
	}

	// dp(index, sum) 表示nums[0]~nums[index]和为sum共有多少种方法
	var dp func(index, sum int) int
	dp = func(index, sum int) int {
		if cache[index][sum+2000] >= 0 {
			return cache[index][sum+2000]
		}

		if index == 0 {
			cnt := 0
			if nums[0] == sum {
				cnt++
			}
			if -nums[0] == sum {
				cnt++
			}
			cache[index][sum+2000] = cnt
			return cache[index][sum+2000]
		}

		cache[index][sum+2000] = dp(index-1, sum+nums[index]) + dp(index-1, sum-nums[index])
		return cache[index][sum+2000]
	}

	return dp(len(nums)-1, target)
}
