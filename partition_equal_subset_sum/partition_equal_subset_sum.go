package partition_equal_subset_sum

// https://leetcode-cn.com/problems/partition-equal-subset-sum/

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 == 1 {
		return false
	}

	cache := make([][]bool, len(nums))
	mark := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		cache[i] = make([]bool, sum/2+1)
		mark[i] = make([]bool, sum/2+1)
	}

	// dp(index, sum)表示nums[0]~nums[index]是否有和为sum的子集
	var dp func(index, sum int) bool
	dp = func(index, sum int) (ret bool) {
		defer func() {
			cache[index][sum] = ret
			mark[index][sum] = true
		}()

		if mark[index][sum] {
			return cache[index][sum]
		}

		if index == 0 {
			return nums[index] == sum
		}

		if nums[index] > sum {
			return dp(index-1, sum)
		}

		return dp(index-1, sum) || dp(index-1, sum-nums[index])
	}

	return dp(len(nums)-1, sum/2)
}
