package daily

// 2643. 一最多的行
func rowAndMaximumOnes(mat [][]int) []int {
	ans := []int{0, 0}
	n := len(mat)
	m := len(mat[0])
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < m; j++ {
			cnt += mat[i][j]
		}
		if cnt > ans[1] {
			ans[0] = i
			ans[1] = cnt
		}
	}
	return ans
}
