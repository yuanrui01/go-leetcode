package daily

// 2109. 向字符串添加空格
func addSpaces(s string, spaces []int) string {
	ans := make([]byte, 0, len(s)+len(spaces))
	j := 0
	for i, c := range s {
		if j < len(spaces) && spaces[j] == i {
			ans = append(ans, ' ')
			j++
		}
		ans = append(ans, byte(c))
	}
	return string(ans)
}
