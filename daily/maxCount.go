package daily

import "math"

// 598. Range Addition II
func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	x, y := math.MaxInt32, math.MaxInt32
	for _, op := range ops {
		x = min(x, op[0])
		y = min(y, op[1])
	}
	return x * y
}
