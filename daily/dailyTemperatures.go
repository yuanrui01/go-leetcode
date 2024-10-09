package daily

// 739. 每日温度
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	var st []int
	for i, t := range temperatures {
		for len(st) > 0 && temperatures[st[len(st)-1]] < t {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ans[j] = i - j
		}
		st = append(st, i)
	}
	return ans
}
