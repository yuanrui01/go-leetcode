package daily

import "strconv"

// 2269. 找到一个数字的 K 美丽值
func divisorSubstrings(num int, k int) int {
	s := strconv.Itoa(num)
	n := len(s)
	ans := 0
	for i := 0; i <= n-k; i++ {
		sub, _ := strconv.Atoi(s[i : i+k])
		if sub != 0 && num%sub == 0 {
			ans++
		}
	}
	return ans
}
