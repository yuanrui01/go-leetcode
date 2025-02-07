package daily

// 63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[n-1][m-1] == 1 {
		return 0
	}
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
	}
	memo[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > 0 && obstacleGrid[i-1][j] != 1 {
				memo[i][j] += memo[i-1][j]
			}
			if j > 0 && obstacleGrid[i][j-1] != 1 {
				memo[i][j] += memo[i][j-1]
			}
		}
	}
	return memo[n-1][m-1]
}
