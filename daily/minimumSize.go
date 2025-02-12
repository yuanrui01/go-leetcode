package daily

import "slices"

// 1760. 袋子里最少数目的球
func minimumSize(nums []int, maxOperations int) int {
	left, right := 1, slices.Max(nums)
	var check func(int) bool
	check = func(m int) bool {
		cnt := 0
		for _, num := range nums {
			cnt += (num - 1) / m
		}
		return cnt <= maxOperations
	}
	for left <= right {
		mid := left + (right-left)/2
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
