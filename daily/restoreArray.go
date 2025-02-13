package daily

// 1743. 从相邻元素对还原数组
func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs)
	vis := make(map[int]bool)
	ans := make([]int, n+1)
	adj := make(map[int][]int, n)

	// 构建邻接表
	for _, pair := range adjacentPairs {
		adj[pair[0]] = append(adj[pair[0]], pair[1])
		adj[pair[1]] = append(adj[pair[1]], pair[0])
	}

	// 找到起始点
	var start int
	for k, v := range adj {
		if len(v) == 1 {
			start = k
			break
		}
	}

	// 构造结果
	ans[0] = start
	vis[start] = true
	for i := 0; i < n; i++ {
		for _, v := range adj[ans[i]] {
			if !vis[v] {
				vis[v] = true
				ans[i+1] = v
				break
			}
		}
	}
	return ans
}
