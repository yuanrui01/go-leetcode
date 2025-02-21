package daily

// 3427. 变长子数组求和
func subarraySum(nums []int) int {
	s := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		s[i+1] = s[i] + nums[i]
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += s[i+1] - s[max(1, i+1-nums[i])-1]
	}
	return ans
}
