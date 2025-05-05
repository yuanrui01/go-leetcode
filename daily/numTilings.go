package daily

//const mod = 1_000_000_007

func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	f := make([]int, n+1)
	f[0], f[1], f[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		f[i] = (2*f[i-1] + f[i-3]) % mod
	}
	return f[n]
}
