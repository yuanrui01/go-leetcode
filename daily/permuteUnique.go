package daily

import "slices"

// 47. 全排列 II
func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, len(nums))
	onPath := make([]bool, len(nums))
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) {
			ans = append(ans, slices.Clone(path))
			return
		}
		for j := 0; j < len(nums); j++ {
			if onPath[j] || j > 0 && nums[j-1] == nums[j] && !onPath[j-1] {
				continue
			}
			onPath[j] = true
			path[i] = nums[j]
			dfs(i + 1)
			onPath[j] = false
		}
	}
	slices.Sort(nums)
	dfs(0)
	return ans
}
