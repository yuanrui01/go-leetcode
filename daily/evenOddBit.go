package daily

// 2595. 奇偶位数
func evenOddBit(n int) []int {
	a := make([]int, 2)
	for i := 0; n != 0; i++ {
		a[i&1] += n & 1
		n >>= 1
	}
	return a
}
