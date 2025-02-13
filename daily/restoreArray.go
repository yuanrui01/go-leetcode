package daily

// 1743. 从相邻元素对还原数组
func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs)
	ans := make([]int, n+1)
	adjMap := make(map[int][]int, n)

	// 构建邻接表
	for _, pair := range adjacentPairs {
		adjMap[pair[0]] = append(adjMap[pair[0]], pair[1])
		adjMap[pair[1]] = append(adjMap[pair[1]], pair[0])
	}

	// 找到起始点
	var start int
	for k, v := range adjMap {
		if len(v) == 1 {
			start = k
			break
		}
	}

	// 构造结果
	ans[0] = start
	for i := 0; i < n; i++ {
		adj := adjMap[ans[i]]
		delete(adjMap, ans[i])
		for _, v := range adj {
			if adjMap[v] != nil {
				ans[i+1] = v
				break
			}
		}
	}
	return ans
}
