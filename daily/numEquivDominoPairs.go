package daily

// 1128. 等价多米诺骨牌对的数量
func numEquivDominoPairs(dominoes [][]int) int {
	flag := make([][10]bool, 10)
	count := make([][10]int, 10)
	ans := 0
	for _, v := range dominoes {
		count[v[0]][v[1]]++
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			c := 0
			if !flag[i][j] {
				flag[i][j] = true
				c += count[i][j]
			}
			if !flag[j][i] {
				flag[j][i] = true
				c += count[j][i]
			}
			if c > 1 {
				ans += c * (c - 1) / 2
			}
		}
	}
	return ans
}
