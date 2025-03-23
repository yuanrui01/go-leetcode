package daily

import "strings"

// 2255. 统计是给定字符串前缀的字符串数目
func countPrefixes(words []string, s string) (ans int) {
	for _, word := range words {
		if strings.HasPrefix(s, word) {
			ans++
		}
	}
	return
}
