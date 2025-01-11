package daily

// 3270. 求出数字答案
func generateKey(num1 int, num2 int, num3 int) int {
	key := 0
	factor := 1
	for i := 0; i < 4; i++ {
		i1 := num1 % 10
		i2 := num2 % 10
		i3 := num3 % 10
		key += factor * min(i1, i2, i3)
		factor *= 10
		num1 /= 10
		num2 /= 10
		num3 /= 10
	}
	return key
}
