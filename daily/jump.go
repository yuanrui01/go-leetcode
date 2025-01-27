package daily

import "math"

// 45. 跳跃游戏 II
func jump(nums []int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == n-1 {
			return 0
		}
		if memo[i] != -1 {
			return memo[i]
		}
		if nums[i] == 0 {
			return math.MaxInt32
		}
		mn := math.MaxInt32
		limit := min(nums[i], n-i-1)
		for s := 1; s <= limit; s++ {
			mn = min(mn, dfs(i+s))
		}
		if mn != math.MaxInt32 {
			mn++
		}
		memo[i] = mn
		return mn
	}
	return dfs(0)
}
