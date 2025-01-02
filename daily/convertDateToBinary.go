package daily

import (
	"strconv"
	"strings"
)

// 3280. 将日期转换为二进制表示
func convertDateToBinary(date string) string {
	a := strings.Split(date, " ")
	for i := range a {
		x, _ := strconv.Atoi(a[i])
		a[i] = strconv.FormatUint(uint64(x), 2)
	}
	return strings.Join(a, "-")
}
