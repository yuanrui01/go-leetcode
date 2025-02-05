package daily

import (
	"slices"
	"sort"
)

// 90. 子集 II
func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0, 1<<n)
	path := make([]int, 0, n)
	sort.Ints(nums)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		// 选
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]

		// 不选
		j := i + 1
		for j < n && nums[j] == nums[i] {
			j++
		}
		dfs(j)
	}
	dfs(0)
	return ans
}
