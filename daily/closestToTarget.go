package daily

import "math"

// 1521. 找到最接近目标值的函数值
func closestToTarget(arr []int, target int) int {
	ans := math.MaxInt32
	for i, v := range arr {
		x := v
		ans = min(ans, abs(x-target))
		for j := i - 1; j >= 0 && ((arr[j] & x) != arr[j]); j-- {
			arr[j] &= x
			ans = min(ans, abs(arr[j]-target))
		}
	}
	return ans
}
