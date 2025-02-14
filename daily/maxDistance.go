package daily

import "sort"

// 1552. 两球之间的磁力
func maxDistance(pos []int, m int) int {
	n := len(pos)
	sort.Ints(pos)
	l, r := 1, (pos[n-1]-pos[0])/(m-1)
	var check func(int) bool
	check = func(mid int) bool {
		pre := pos[0]
		cnt := m - 1
		for i := 1; i < n; i++ {
			if pos[i]-pre >= mid {
				cnt--
				pre = pos[i]
			}
			if cnt == 0 {
				return true
			}
		}
		return false
	}
	for l <= r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}
