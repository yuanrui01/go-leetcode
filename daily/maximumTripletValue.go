package daily

// 2873. 有序三元组中的最大值 I
func maximumTripletValue(nums []int) int64 {
	ans := int64(0)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			if nums[i] > nums[j] {
				for k := j + 1; k < n; k++ {
					ans = max(ans, int64(nums[i]-nums[j])*int64(nums[k]))
				}
			}
		}
	}
	return ans
}
