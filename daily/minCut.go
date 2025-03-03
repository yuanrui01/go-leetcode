package daily

import "math"

// 132. 分割回文串 II
func minCut(s string) int {
	n := len(s)
	palMemo := make([][]int8, n)
	for i := range palMemo {
		palMemo[i] = make([]int8, n)
		for j := range palMemo[i] {
			palMemo[i][j] = -1
		}
	}

	var isPalidrome func(int, int) bool
	isPalidrome = func(l, r int) bool {
		if l >= r {
			return true
		}
		p := &palMemo[l][r]
		if *p != -1 {
			return *p == 1
		}
		res := s[l] == s[r] && isPalidrome(l+1, r-1)
		if res {
			*p = 1
		} else {
			*p = 0
		}
		return res
	}
	dfsMemo := make([]int, n)
	for i := range dfsMemo {
		dfsMemo[i] = -1
	}
	var dfs func(int) int
	dfs = func(r int) int {
		if isPalidrome(0, r) {
			return 0
		}
		p := &dfsMemo[r]
		if *p != -1 {
			return *p
		}
		res := math.MaxInt32
		for l := 1; l <= r; l++ {
			if isPalidrome(l, r) {
				res = min(res, dfs(l-1)+1)
			}
		}
		*p = res
		return res
	}
	return dfs(n - 1)
}
