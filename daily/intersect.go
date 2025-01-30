package daily

// 350. 两个数组的交集 II
func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	mode1 := make([]int, 1001)
	for _, v := range nums1 {
		mode1[v]++
	}
	for _, v := range nums2 {
		if mode1[v] > 0 {
			res = append(res, v)
			mode1[v]--
		}
	}
	return res
}
