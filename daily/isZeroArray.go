package daily

// 3355. 零数组变换 I
func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	diff := make([]int, n+1)
	for _, q := range queries {
		diff[q[0]]++
		diff[q[1]+1]--
	}
	for i := 0; i < n; i++ {
		if nums[i] > diff[i] {
			return false
		}
		diff[i+1] += diff[i]
	}
	return true
}
