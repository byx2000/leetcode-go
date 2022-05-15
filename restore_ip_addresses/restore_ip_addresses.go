package restore_ip_addresses

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/restore-ip-addresses/

func restoreIpAddresses(s string) []string {
	// 判断s是不是合法的ip数字
	valid := func(s string) bool {
		if len(s) == 0 {
			return false
		}
		if s[0] == '0' && len(s) > 1 {
			return false
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if i > 255 {
			return false
		}
		return true
	}

	// 将s分割成s[:i]、s[i:j]、s[j:k]、s[k:]四段
	var result []string
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			for k := j; k < len(s); k++ {
				if valid(s[:i]) && valid(s[i:j]) && valid(s[j:k]) && valid(s[k:]) {
					result = append(result, fmt.Sprintf("%s.%s.%s.%s", s[:i], s[i:j], s[j:k], s[k:]))
				}
			}
		}
	}

	return result
}
