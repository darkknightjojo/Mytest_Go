package Leetcode

import (
	"sort"
)

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	lastStr := ""
	answer := ""
	if len(strs) == 1 {
		return strs[0]
	}
	for _, s := range strs {
		if s == "" {
			return ""
		}
		temp := 0
		for j := 0; j < len(s) && j < len(lastStr); j++ {
			if s[j] == lastStr[j] {
				temp++
			} else {
				break
			}
		}
		if temp == 0 && lastStr != "" {
			return ""
		}
		if answer == "" || len(answer) > temp {
			answer = s[:temp]
		}
		lastStr = s
	}
	return answer
}
