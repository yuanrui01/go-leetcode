package daily

import "slices"

// 3375. 使数组的值全部为 K 的最少操作次数
func minOperations3375(nums []int, k int) int {
	mn := slices.Min(nums)
	if k > mn {
		return -1
	}

	set := map[int]struct{}{}
	for _, x := range nums {
		set[x] = struct{}{}
	}
	if k == mn {
		return len(set) - 1
	}
	return len(set)
}
