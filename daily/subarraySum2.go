package daily

// SubarraySum2 560. 和为 K 的子数组
func SubarraySum2(nums []int, k int) int {
	m := map[int]int{}
	m[0] = 1
	ans := 0
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
		v, e := m[s-k]
		if e {
			ans += v
		}
		m[s]++
	}
	return ans
}
