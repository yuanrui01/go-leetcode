package daily

import "slices"

// Partition 131. 分割回文串
func Partition(s string) [][]string {
	n := len(s)
	ans := [][]string{}
	path := []string{}
	var isPalindrome func(s string, l int, r int) bool
	isPalindrome = func(s string, l int, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		if i == n-1 {
			if isPalindrome(s, j, i) {
				path = append(path, s[j:i+1])
				ans = append(ans, slices.Clone(path))
				path = path[:len(path)-1]
			}
			return
		}
		if isPalindrome(s, j, i) {
			path = append(path, s[j:i+1])
			dfs(i+1, i+1)
			path = path[:len(path)-1]
		}
		dfs(i+1, j)
	}
	dfs(0, 0)
	return ans
}
