package daily

// 3396. 使数组元素互不相同所需的最少操作次数
func minimumOperations(nums []int) int {
	mark := make([]bool, 101)
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		if mark[nums[i]] {
			return i/3 + 1
		}
		mark[nums[i]] = true
	}
	return 0
}
