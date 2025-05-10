package daily

// 2918. 数组的最小相等和
func minSum(nums1 []int, nums2 []int) int64 {
	rs1 := make([]int, 2)
	rs2 := make([]int, 2)
	for _, x := range nums1 {
		rs1[0] += x
		if x == 0 {
			rs1[1]++
		}
	}
	for _, x := range nums2 {
		rs2[0] += x
		if x == 0 {
			rs2[1]++
		}
	}
	minS1 := rs1[0] + rs1[1]
	minS2 := rs2[0] + rs2[1]
	if minS1 == minS2 {
		return int64(minS1)
	}
	if minS1 > minS2 && rs2[1] > 0 {
		return int64(minS1)
	}
	if minS2 > minS1 && rs1[1] > 0 {
		return int64(minS2)
	}
	return -1
}
