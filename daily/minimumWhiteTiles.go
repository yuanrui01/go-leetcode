package daily

// 2209. 用地毯覆盖后的最少白色砖块
func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	m := len(floor)
	dp := make([][]int, numCarpets+1)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0] = int(floor[0] - '0')
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + int(floor[i]-'0')
	}
	for i := 1; i <= numCarpets; i++ {
		for j := i * carpetLen; j < m; j++ {
			dp[i][j] = min(dp[i-1][j-carpetLen], dp[i][j-1]+int(floor[j]-'0'))
		}
	}
	return dp[numCarpets][m-1]
}
