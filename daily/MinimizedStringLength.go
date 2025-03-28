package daily

// 2716. 最小化字符串长度
func minimizedStringLength(s string) int {
	ans := 0
	mask := 0
	n := len(s)
	for i := 0; i < n; i++ {
		mask |= 1 << (s[i] - 'a')
	}
	for i := 0; i < 26; i++ {
		if mask&(1<<uint(i)) != 0 {
			ans++
		}
	}
	return ans
}
