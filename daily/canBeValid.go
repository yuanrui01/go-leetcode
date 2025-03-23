package daily

// 2116. 判断一个括号字符串是否有效
func canBeValid(s string, locked string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	mx, mn := 0, 0
	for i := 0; i < n; i++ {
		if locked[i] == '1' {
			d := -1
			if s[i] == '(' {
				d = 1
			}
			mx += d
			if mx < 0 {
				return false
			}
			mn += d
		} else {
			mx++
			mn--
		}
		if mn < 0 {
			mn = 1
		}
	}
	return mn == 0
}
