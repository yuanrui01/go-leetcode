package daily

// 2614. 对角线上的质数
func diagonalPrime(nums [][]int) int {
	var isPrime func(n int) bool
	isPrime = func(n int) bool {
		if n <= 1 {
			return false
		}
		if n <= 3 {
			return true
		}
		if n%2 == 0 || n%3 == 0 {
			return false
		}
		for i := 5; i*i <= n; i += 6 {
			if n%i == 0 || n%(i+2) == 0 {
				return false
			}
		}
		return true
	}

	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		x1 := nums[i][i]
		x2 := nums[i][n-i-1]
		if x1 > ans && isPrime(x1) {
			ans = x1
		}
		if x2 > ans && isPrime(x2) {
			ans = x2
		}
	}
	return ans
}
