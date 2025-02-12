package daily

import (
	"slices"
	"sort"
)

// 1760. 袋子里最少数目的球
func minimumSize(nums []int, maxOperations int) int {
	left, right := 1, slices.Max(nums)
	return left + sort.Search(right-left, func(m int) bool {
		m += left
		cnt := 0
		for _, num := range nums {
			cnt += (num - 1) / m
		}
		return cnt <= maxOperations
	})
}
