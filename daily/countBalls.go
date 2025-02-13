package daily

// 1742. 盒子中小球的最大数量
func countBalls(lowLimit int, highLimit int) int {
	count := make([]int, 46)
	var digitSum func(int) int
	digitSum = func(num int) int {
		sum := 0
		for num != 0 {
			sum += num % 10
			num /= 10
		}
		return sum
	}
	for i := lowLimit; i <= highLimit; i++ {
		count[digitSum(i)]++
	}
	mx := 0
	for _, v := range count {
		mx = max(mx, v)
	}
	return mx
}
