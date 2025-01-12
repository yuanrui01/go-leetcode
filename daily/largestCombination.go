package daily

import "slices"

// 2275. 按位与结果大于零的最长组合
func largestCombination(candidates []int) int {
	cnt := [24]int{}
	for _, v := range candidates {
		num := v
		for i := 0; num != 0; i++ {
			cnt[i] += num & 1
			num >>= 1
		}
	}
	return slices.Max(cnt[:])
}

//func largestCombination_bak(candidates []int) int {
//	cnt := make([]int, 24)
//	ans := 0
//	for _, v := range candidates {
//		k := 0
//		num := v
//		for num != 0 {
//			cnt[k] += num & 1
//			k++
//			num >>= 1
//		}
//	}
//	for _, v := range cnt {
//		ans = max(ans, v)
//	}
//	return ans
//}
