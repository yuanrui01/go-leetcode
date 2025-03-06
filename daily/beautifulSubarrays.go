package daily

// 2588. 统计美丽子数组数目
func beautifulSubarrays(nums []int) int64 {
	ans := int64(0)
	s := 0
	m := map[int]int{}
	m[0] = 1
	for _, v := range nums {
		s ^= v
		v, e := m[s]
		if e {
			ans += int64(v)
		}
		m[s]++
	}
	return ans
}
