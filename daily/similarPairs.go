package daily

// 2506. 统计相似字符串对的数目
func similarPairs(words []string) int {
	n := len(words)
	m := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < len(words[i]); j++ {
			m[i] |= (1 << uint(words[i][j]-'a'))
		}
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if m[i] == m[j] {
				ans++
			}
		}
	}
	return ans
}
