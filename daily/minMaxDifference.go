package daily

import "strconv"

// 2566. 替换一个数字后的最大差值
func minMaxDifference(num int) int {
	s := strconv.Itoa(num)
	firstC := byte(s[0])
	firstNot9 := byte('.')
	for i := 0; i < len(s); i++ {
		if s[i] != '9' {
			firstNot9 = s[i]
			break
		}
	}
	mx, mn := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == firstC {
			mn = mn * 10
		} else {
			mn = mn*10 + int(s[i]-'0')
		}
		if s[i] == firstNot9 {
			mx = mx*10 + 9
		} else {
			mx = mx*10 + int(s[i]-'0')
		}
	}
	return mx - mn
}
