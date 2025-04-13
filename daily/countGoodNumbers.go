package daily

// 1922. 统计好数字的数目
func countGoodNumbers(n int64) int {
	cache := make(map[int]int)
	var dfs func(l int) int
	dfs = func(l int) int {
		if l == 1 {
			return 5
		}
		if l == 2 {
			return 20
		}
		if v, ok := cache[l]; ok {
			return v
		}
		l1 := l / 2
		if l1%2 != 0 {
			l1++
		}
		l2 := l - l1
		a1 := dfs(l1)
		a2 := dfs(l2)
		a3 := (a1 * a2) % mod
		cache[l] = a3
		return a3
	}
	return dfs(int(n))
}
