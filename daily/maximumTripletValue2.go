package daily

// 2874. 有序三元组中的最大值 II
func maximumTripletValue2(nums []int) int64 {
	ans := int64(0)
	n := len(nums)
	rm := make([]int, n-2)
	rm[n-3] = nums[n-1]
	lm := make([]int, n-2)
	lm[0] = nums[0]
	for i := n - 2; i > 1; i-- {
		rm[i-2] = max(rm[i-1], nums[i])
	}
	for i := 1; i < n-2; i++ {
		lm[i] = max(lm[i-1], nums[i])
	}
	for i := 1; i < n-1; i++ {
		ans = max(ans, int64(lm[i-1]-nums[i])*int64(rm[i-1]))
	}
	return ans
}
