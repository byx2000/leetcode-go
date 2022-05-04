package letter_combinations_of_a_phone_number

// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := map[uint8]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string

	var dfs func(index int, word string)
	dfs = func(index int, word string) {
		if index == len(digits) {
			result = append(result, word)
			return
		}

		for _, c := range m[digits[index]] {
			dfs(index+1, word+string(c))
		}
	}

	dfs(0, "")
	return result
}
