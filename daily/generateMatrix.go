package daily

// 59. 螺旋矩阵 II
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	var f func(num int, l int, base int)
	f = func(num int, l int, base int) {
		for j := 0; j < l; j++ {
			ans[base][base+j] = num
			num++
		}
		for i := 1; i < l; i++ {
			ans[base+i][base+l-1] = num
			num++
		}
		for j := l - 2; j >= 0; j-- {
			ans[base+l-1][base+j] = num
			num++
		}
		for i := l - 2; i >= 1; i-- {
			ans[base+i][base] = num
			num++
		}
		if l-2 > 0 {
			f(num, l-2, base+1)
		}
	}
	f(1, n, 0)
	return ans
}
