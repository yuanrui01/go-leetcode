package daily

// 2275. 按位与结果大于零的最长组合
func largestCombination(candidates []int) int {
	cnt := make([]int, 24)
	ans := 0
	for _, v := range candidates {
		k := 0
		num := v
		for num != 0 {
			cnt[k] += num & 1
			k++
			num >>= 1
		}
	}
	for _, v := range cnt {
		ans = max(ans, v)
	}
	return ans
}
