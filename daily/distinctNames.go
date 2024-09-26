package daily

// 2306. 公司命名
func distinctNames(ideas []string) (ans int64) {
	size := [26]int{}
	intersection := [26][26]int{}
	groups := map[string]int{}
	for _, s := range ideas {
		b := s[0] - 'a'
		size[b]++
		suffix := s[1:]
		mask := groups[suffix]
		groups[suffix] = mask | 1<<b
		for a := 0; a < 26; a++ {
			if mask>>a&1 > 0 {
				intersection[b][a]++
				intersection[a][b]++
			}
		}
	}
	for i := 1; i < 26; i++ {
		for j := 0; j < i; j++ {
			m := intersection[i][j]
			ans += int64(size[i]-m) * int64(size[j]-m)
		}
	}
	return ans * 2
}
