package daily

// 1399. 统计最大组的数目
func countLargestGroup(n int) int {
	m := make(map[int]int)
	ans := 0
	mx := 0
	for i := 1; i <= n; i++ {
		s := 0
		for t := i; t != 0; t /= 10 {
			s += t % 10
		}
		m[s]++
		if m[s] > mx {
			mx = m[s]
			ans = 1
		} else if m[s] == mx {
			ans++
		}
	}
	return ans
}
