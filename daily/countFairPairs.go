package daily

import (
	"slices"
	"sort"
)

// 2563. 统计公平数对的数目
func countFairPairs(nums []int, lower int, upper int) (ans int64) {
	slices.Sort(nums)
	for j, x := range nums {
		r := sort.SearchInts(nums[:j], upper-x+1)
		l := sort.SearchInts(nums[:r], lower-x)
		ans += int64(r - l)
	}
	return
}
