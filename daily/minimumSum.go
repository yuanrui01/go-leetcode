package daily

// 2829. k-avoiding 数组的最小总和
func minimumSum(n int, k int) int {
	if k/2 > n {
		return n * (n + 1) / 2
	}
	i1 := k / 2
	ans := i1 * (i1 + 1) / 2
	i2 := n - i1
	if i2 > 0 {
		ans += i2 * (2*k + i2 - 1) / 2
	}
	return ans
}
