package daily

// 1706. 球会落何处
func findBall(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])
	ans := make([]int, n)
	var getFinalPos func(pj int) int
	getFinalPos = func(pj int) int {
		for i := 0; i < m; i++ {
			if grid[i][pj] == 1 {
				if pj == n-1 || grid[i][pj+1] == -1 {
					return -1
				} else {
					pj++
				}
			} else {
				if pj == 0 || grid[i][pj-1] == 1 {
					return -1
				} else {
					pj--
				}
			}
		}
		return pj
	}
	for j := 0; j < n; j++ {
		ans[j] = getFinalPos(j)
	}
	return ans
}
