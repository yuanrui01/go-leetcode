package daily

// 2680. 最大或值
func maximumOr(nums []int, k int) int64 {
	n := len(nums)
	suf := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1] | nums[i+1]
	}
	ans := 0
	pre := 0
	for i := 0; i < n; i++ {
		ans = max(ans, pre|nums[i]<<k|suf[i])
		pre |= nums[i]
	}
	return int64(ans)
}
