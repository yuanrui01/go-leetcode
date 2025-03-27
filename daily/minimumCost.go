package daily

// 2712. 使所有字符相等的最小成本
func minimumCost(s string) int64 {
	n := len(s)
	if n == 1 {
		return 0
	}
	ans0, ans1 := 0, 0
	b0, b1 := false, false

	for s1 := n/2 - 1; s1 >= 0; s1-- {
		if s[s1] == '0' && !b1 || s[s1] == '1' && b1 {
			ans1 += (s1 + 1)
			b1 = !b1
		}
		if s[s1] == '1' && !b0 || s[s1] == '0' && b0 {
			ans0 += (s1 + 1)
			b0 = !b0
		}
	}
	b1 = false
	b0 = false
	for s2 := n / 2; s2 < n; s2++ {
		if s[s2] == '0' && !b1 || s[s2] == '1' && b1 {
			ans1 += (n - s2)
			b1 = !b1
		}
		if s[s2] == '1' && !b0 || s[s2] == '0' && b0 {
			ans0 += (n - s2)
			b0 = !b0
		}
	}
	return int64(min(ans0, ans1))
}
