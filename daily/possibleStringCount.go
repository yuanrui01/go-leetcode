package daily

// 3330. 找到初始输入字符串 I
func possibleStringCount(word string) int {
	ans, i, n := 1, 0, len(word)
	for i < n {
		cnt, x := 1, word[i]
		for i+1 < n && word[i+1] == x {
			cnt++
			i++
		}
		ans += cnt - 1
		i++
	}
	return ans
}
