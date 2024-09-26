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

	for a := 1; a < 26; a++ {
		for b := 0; b < a; b++ {
			m := intersection[a][b]
			ans += int64(size[a]-m) * int64(size[b]-m)
		}
	}

	return ans * 2
}
