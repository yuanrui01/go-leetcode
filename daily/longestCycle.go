package daily

// 2360. 图中的最长环
func longestCycle(edges []int) int {
	ans := -1
	n := len(edges)
	visTime := make([]int, n)
	curTime := 1
	for i := 0; i < n; i++ {
		x := i
		startTime := curTime
		for x != -1 && visTime[x] == 0 {
			visTime[x] = curTime
			x = edges[x]
			curTime++
		}
		if x != -1 && visTime[x] > startTime {
			ans = max(ans, visTime[x]-startTime)
		}
	}
	return ans
}
