package daily

// 386. 字典序排数
func lexicalOrder(mx int) (res []int) {
	var dfs func(int)
	dfs = func(n int) {
		if n > mx {
			return
		}
		res = append(res, n)
		nn := 10 * n
		for i := 0; i < 10; i++ {
			dfs(nn + i)
		}
	}
	for i := 1; i < 10; i++ {
		dfs(i)
	}
	return
}
