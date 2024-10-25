package dp

// 746. 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	f0, f1 := 0, 0
	for i := 2; i <= n; i++ {
		t := f1
		f1 = min(f0+cost[i-2], f1+cost[i-1])
		f0 = t
	}
	return f1
}
