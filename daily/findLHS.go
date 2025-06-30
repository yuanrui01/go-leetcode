package daily

import "sort"

// 594. 最长和谐子序列
func findLHS(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	if n == 1 || nums[0] == nums[n-1] {
		return 0
	}
	ans := 0
	l, r := 0, 1
	for l < n {
		for r < n && nums[r]-nums[l] <= 1 {
			r++
		}
		if nums[l] != nums[r-1] {
			ans = max(ans, r-l)
		}
		l++
	}
	return ans
}
