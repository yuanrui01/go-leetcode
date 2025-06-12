package daily

// 3423. 循环数组中相邻元素的最大差值
func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	ans := abs(nums[0] - nums[n-1])
	for i := 1; i < n; i++ {
		ans = max(ans, abs(nums[i]-nums[i-1]))
	}
	return ans
}

// func abs(v int) int {
// 	if v < 0 {
// 		return -v
// 	}
// 	return v
// }
