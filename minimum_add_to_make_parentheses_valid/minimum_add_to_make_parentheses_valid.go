package minimum_add_to_make_parentheses_valid

// https://leetcode.cn/problems/minimum-add-to-make-parentheses-valid/

func minAddToMakeValid(s string) int {
	var stack []int32

	for _, c := range s {
		if c == '(' {
			stack = append(stack, c)
		} else {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[0 : len(stack)-1]
			} else {
				stack = append(stack, c)
			}
		}
	}

	return len(stack)
}
