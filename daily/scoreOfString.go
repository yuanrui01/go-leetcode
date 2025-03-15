package daily

// 3110. 字符串的分数
func scoreOfString(S string) int {
	s := []byte(S)
	ans := 0
	for i := 1; i < len(s); i++ {
		ans += abs(int(s[i]) - int(s[i-1]))
	}
	return ans
}
