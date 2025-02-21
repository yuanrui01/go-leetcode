package daily

// 2209. 用地毯覆盖后的最少白色砖块
func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	m := len(floor)
	memo := make([][]int, numCarpets+1)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if j < carpetLen*i {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		res = dfs(i, j-1) + int(floor[j]-'0')
		if i > 0 {
			res = min(res, dfs(i-1, j-carpetLen))
		}
		return
	}
	return dfs(numCarpets, m-1)
}
