package daily

// 416. 分割等和子集
func canPartition(nums []int) bool {
	n := len(nums)
	s := 0
	for _, num := range nums {
		s += num
	}
	if s%2 != 0 {
		return false
	}
	m := s/2 + 1
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, m)
		for j := 0; j < m; j++ {
			memo[i][j] = -1
		}
	}
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i < 0 {
			return j == 0
		}
		if memo[i][j] != -1 {
			return memo[i][j] == 1
		}
		res := j >= nums[i] && dfs(i-1, j-nums[i]) || dfs(i-1, j)
		if res {
			memo[i][j] = 1
		} else {
			memo[i][j] = 0
		}
		return res
	}
	return dfs(n-1, s/2)
}
