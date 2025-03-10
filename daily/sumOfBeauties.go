package daily

// 2012. 数组美丽值求和
func sumOfBeauties(nums []int) int {
	n := len(nums)
	mx := nums[0]
	ans := 0
	mn := make([]int, n)
	mn[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		mn[i] = min(mn[i+1], nums[i])
	}
	for i := 1; i < n-1; i++ {
		if nums[i] > mx && nums[i] < mn[i+1] {
			ans += 2
		} else if nums[i] > nums[i-1] && nums[i] < nums[i+1] {
			ans++
		}
		mx = max(mx, nums[i])
	}
	return ans
}
