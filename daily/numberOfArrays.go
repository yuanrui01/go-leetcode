package daily

// 2145. 统计隐藏数组数目
func numberOfArrays(differences []int, lower int, upper int) int {
	var s, minS, maxS int
	for _, x := range differences {
		s += x
		maxS = max(maxS, s)
		minS = min(minS, s)
	}
	return max(0, upper-lower-maxS+minS+1)
}
