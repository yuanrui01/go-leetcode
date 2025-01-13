package daily

// 2270. 分割数组的方案数
func waysToSplitArray(nums []int) int {
	sum := int64(0)
	for _, v := range nums {
		sum += int64(v)
	}
	ans := 0
	sufSum := int64(0)
	for i := len(nums) - 1; i > 0; i-- {
		sufSum += int64(nums[i])
		if sum >= 2*sufSum {
			ans++
		}
	}
	return ans
}
