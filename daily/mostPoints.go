package daily

// 2140. 解决智力问题
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		p := questions[i][0]
		b := questions[i][1]
		dp[i] = p
		if i+b+1 < n {
			dp[i] += dp[i+b+1]
		}
		if i+1 < n {
			dp[i] = max(dp[i], dp[i+1])
		}
	}
	return int64(dp[0])
}
