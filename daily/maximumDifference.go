package daily

// 2016. 增量元素之间的最大差值
func maximumDifference(nums []int) int {
	mn, ans := nums[0], -1

	for _, x := range nums {
		if x > mn {
			ans = max(ans, x-mn)
		}
		mn = min(mn, x)
	}
	return ans
}
