package daily

// 1957. 删除字符使字符串变好
func makeFancyString(s string) string {
    ans := []byte{}
	i := 0
	n := len(s)
	for i < n {
		ch := s[i]
		cnt := 1
		ans = append(ans, byte(ch))
		for i+1 < n && s[i+1] == ch {
			i++
			cnt++
			if cnt < 3 {
				ans = append(ans, byte(ch))
			}
		}
		i++
	}
	return string(ans)
}