package daily

// 3065. 超过阈值的最少操作数 I
func minOperations(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			ans++
		}
	}
	return ans
}
