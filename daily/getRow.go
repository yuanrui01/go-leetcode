package daily

// 119. 杨辉三角 II
func getRow(rowIndex int) []int {
	var ans [][]int
	ans = append(ans, []int{1})
	for i := 1; i <= rowIndex; i++ {
		p := ans[i-1]
		l := []int{1}
		for j := 0; j < len(p)-1; j++ {
			l = append(l, p[j]+p[j+1])
		}
		l = append(l, 1)
		ans = append(ans, l)
	}
	return ans[rowIndex]
}
