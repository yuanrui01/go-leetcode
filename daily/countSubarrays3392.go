package daily

// 3392. 统计符合条件长度为 3 的子数组数目
func countSubarrays(nums []int) int {
	n := len(nums)
	ans, i, j, k := 0, 0, 1, 2
	for k < n {
		if 2*(nums[i]+nums[k]) == nums[j] {
			ans++
		}
		i++
		k++
		j++
	}
	return ans
}
