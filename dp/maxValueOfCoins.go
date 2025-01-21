package dp

// 2218. 从栈中取出 K 个硬币的最大面值和
func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, k+1)
	}

	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p != 0 {
			return *p
		}
		defer func() { *p = res }()

		res = dfs(i-1, j)
		v := 0
		for w := range min(j, len(piles[i])) {
			v += piles[i][w]
			res = max(res, dfs(i-1, j-w-1)+v)
		}
		return
	}
	return dfs(n-1, k)
}
