package daily

// 1963. 使字符串平衡的最小交换次数
func minSwaps1963(S string) int {
	s := []byte(S)
	ans, c, j := 0, 0, len(s)-1
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			c++
		} else if c > 0 {
			c--
		} else {
			for s[j] == ']' {
				j--
			}
			s[j] = ']'
			c++
			ans++
		}
	}
	return ans
}
