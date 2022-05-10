package daily_temperatures

// https://leetcode.cn/problems/daily-temperatures/

func dailyTemperatures(temperatures []int) []int {
	var stack []int
	result := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			topIdx := stack[len(stack)-1]
			result[topIdx] = i - topIdx
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}
