package daily

import "slices"

// 541. 反转字符串 II
func reverseStr(S string, k int) string {
	s := []byte(S)
	n := len(s)
	for i := 0; i < n; i += 2 * k {
		slices.Reverse(s[i:min(i+k, n)])
	}
	return string(s)
}
