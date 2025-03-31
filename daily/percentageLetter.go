package daily

// 2278. 字母在字符串中的百分比
func percentageLetter(s string, letter byte) int {
	n := len(s)
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == letter {
			cnt++
		}
	}
	return (100.0 * cnt) / n
}
