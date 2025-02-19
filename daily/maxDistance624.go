package daily

// 624. 数组列表中的最大距离
func maxDistance624(arrays [][]int) int {
	array0 := arrays[0]
	mn, mx := array0[0], array0[len(array0)-1]
	ans := -1
	for i := 1; i < len(arrays); i++ {
		first := arrays[i][0]
		last := arrays[i][len(arrays[i])-1]
		ans = max(ans, abs(mx-first))
		ans = max(ans, abs(mn-last))
		mx = max(mx, last)
		mn = min(mn, first)
	}
	return ans
}
