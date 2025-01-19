package daily

// FindClosestNumber 2239. 找到最接近 0 的数字
func FindClosestNumber(nums []int) int {
	ans := nums[0]
	absA := abs(ans)
	for i := 1; i < len(nums); i++ {
		absN := abs(nums[i])
		if absN < absA || (absN == absA && nums[i] > ans) {
			ans = nums[i]
			absA = abs(ans)
		}
	}
	return ans
}
