package permutation_in_string

// https://leetcode.cn/problems/permutation-in-string/

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	freq1 := make([]int, 26)
	for _, c := range s1 {
		freq1[c-'a']++
	}

	freq2 := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		freq2[s2[i]-'a']++
	}

	match := func() bool {
		for i := 0; i < 26; i++ {
			if freq1[i] != freq2[i] {
				return false
			}
		}
		return true
	}

	if match() {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		freq2[s2[i]-'a']++
		freq2[s2[i-len(s1)]-'a']--
		if match() {
			return true
		}
	}

	return false
}
