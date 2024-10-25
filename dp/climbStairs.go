package dp

// 70. 爬楼梯
func climbStairs(n int) int {
	f0, f1 := 1, 1
	for i := 2; i <= n; i++ {
		t := f1
		f1 = f1 + f0
		f0 = t
	}
	return f1
}
