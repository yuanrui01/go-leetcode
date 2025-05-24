package daily

import "strings"

// 2942. 查找包含给定字符的单词
func findWordsContaining(words []string, x byte) (ans []int) {
	for i, s := range words {
		if strings.IndexByte(s, x) >= 0 {
			ans = append(ans, i)
		}
	}
	return
}
