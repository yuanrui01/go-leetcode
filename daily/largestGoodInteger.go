package daily

import (
	"strings"
)

// LargestGoodInteger 2264. 字符串中最大的 3 位相同数字
func LargestGoodInteger(num string) string {
	ans := uint8(0)
	for i := 2; i < len(num); i++ {
		if num[i] > ans && num[i] == num[i-1] && num[i] == num[i-2] {
			ans = num[i]
		}
	}
	if ans > 0 {
		return strings.Repeat(string(ans), 3)
	}
	return ""
}
