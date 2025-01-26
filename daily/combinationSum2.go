package daily

import (
	"slices"
)

// 40. 组合总和 II
func combinationSum2(candidates []int, target int) (ans [][]int) {
	slices.Sort(candidates)
	n := len(candidates)
	path := []int{}
	var dfs func(i, left int)
	dfs = func(i, left int) {
		if left == 0 {
			ans = append(ans, slices.Clone(path))
			return
		}
		if i == n {
			return
		}
		x := candidates[i]
		if left < x {
			return
		}
		path = append(path, x)
		dfs(i+1, left-x)
		path = path[:len(path)-1]

		i++
		for i < n && candidates[i] == x {
			i++
		}
		dfs(i, left)
	}
	dfs(0, target)
	return ans
}
