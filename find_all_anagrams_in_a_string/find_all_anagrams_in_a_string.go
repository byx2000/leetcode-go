package find_all_anagrams_in_a_string

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	freq1 := make([]int, 26)
	for _, c := range p {
		freq1[c-'a']++
	}

	freq2 := make([]int, 26)
	for i := 0; i < len(p); i++ {
		freq2[s[i]-'a']++
	}

	match := func() bool {
		for i := 0; i < 26; i++ {
			if freq1[i] != freq2[i] {
				return false
			}
		}
		return true
	}

	var result []int

	if match() {
		result = append(result, 0)
	}

	for i := len(p); i < len(s); i++ {
		freq2[s[i-len(p)]-'a']--
		freq2[s[i]-'a']++
		if match() {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}
