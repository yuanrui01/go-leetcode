package daily

// 2302. 统计得分小于 K 的子数组数目
func countSubarrays2302(nums []int, k int64) int64 {
	kk := int(k)
	sum, count, left, right, n := 0, 0, 0, 0, len(nums)

	for right < n {
		sum += nums[right]
		for left <= right && sum*(right-left+1) >= kk {
			count += n - right
			sum -= nums[left]
			left++
		}
		right++
	}
	return int64(n*(n+1)/2 - count)
}
