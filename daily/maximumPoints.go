package daily

// 2920. 收集所有金币可获得的最大积分
func maximumPoints(edges [][]int, coins []int, k int) int {
	n := len(coins)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	vis := make([]bool, n)
	memo := make([][]int, len(coins))
	depth := getDepth(coins)
	for i := range memo {
		memo[i] = make([]int, depth+1)
	}
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	return dfs(0, k, 1, g, coins, vis, memo)
}

func dfs(i int, k int, factor int, g [][]int, coins []int, vis []bool, memo [][]int) int {
	if memo[i][factor] != -1 {
		return memo[i][factor]
	}
	vis[i] = true
	s1, s2 := 0, 0
	for _, v := range g[i] {
		if !vis[v] {
			s1 += dfs(v, k, factor, g, coins, vis, memo)
			if k > 0 && factor < len(memo[0])-1 {
				s2 += dfs(v, k, factor+1, g, coins, vis, memo)
			}
		}
	}
	vis[i] = false
	actV := coins[i] / (1 << (factor - 1))
	memo[i][factor] = max(int(actV-k+s1), int(actV/2+s2))
	return memo[i][factor]
}

func getDepth(coins []int) int {
	mx := 0
	for _, coin := range coins {
		mx = max(mx, coin)
	}
	depth := 1
	for mx/2 != 0 {
		depth++
		mx /= 2
	}
	return depth
}
