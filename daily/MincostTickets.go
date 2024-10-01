package daily

// 983. 最低票价
func mincostTickets(days []int, costs []int) int {
	limit := days[len(days)-1]
	dp := make([]int, limit+1)
	j := 0
	for i := 1; i <= limit; i++ {
		if days[j] == i {
			cost1 := dp[i-1] + costs[0]
			cost2 := dp[max(0, i-7)] + costs[1]
			cost3 := dp[max(0, i-30)] + costs[2]
			dp[i] = min(cost1, cost2, cost3)
			j++
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[limit]
}
