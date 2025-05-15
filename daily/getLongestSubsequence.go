func getLongestSubsequence(words []string, groups []int) []string {
	n := len(groups)

	ans := []string{}
	for i, x := range groups {
		if i == n-1 || x != groups[i+1] {
			ans = append(ans, words[i])
		}
	}
	return ans
}