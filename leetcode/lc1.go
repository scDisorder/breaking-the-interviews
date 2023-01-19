package leetcode

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]

	for i := range strs {
		for strings.Index(strs[i], res) != 0 {
			res = res[:len(res)-1]
			if res == "" {
				return ""
			}
		}
	}

	return res
}
