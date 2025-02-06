package daily

import "slices"

// 46. 全排列
func permute(nums []int) [][]int {
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
			if !onPath[j] {
				onPath[j] = true
				path[i] = nums[j]
				dfs(i + 1)
				onPath[j] = false
			}
		}
	}
	dfs(0)
	return ans
}
