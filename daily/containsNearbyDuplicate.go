package daily

// 219. 存在重复元素 II
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		v, e := m[nums[i]]
		if e && abs(i-v) <= k {
			return true
		}
		m[nums[i]] = i
	}
	return false
}
