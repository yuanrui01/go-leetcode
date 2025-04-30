package daily

// 1295. 统计位数为偶数的数字
func findNumbers(nums []int) int {
	ans := 0
	for _, x := range nums {
		for x >= 100 {
			x /= 100
		}
		if x >= 10 {
			ans += 1
		}
	}
	return ans
}
