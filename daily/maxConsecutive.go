package daily

import "sort"

// 2274. 不含特殊楼层的最大连续楼层数
func maxConsecutive(bottom int, top int, special []int) int {
	sort.Ints(special)
	length := len(special)
	ans := special[0] - bottom
	ans = mMax(ans, top-special[length-1])
	for i := 1; i < length; i++ {
		ans = mMax(ans, special[i]-special[i-1]-1)
	}
	return ans
}

func mMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
