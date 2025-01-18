package daily

func longestCommonPrefix(strs []string) string {
    s0 := strs[0]
	for i := range(s0) {
		for _, str := range(strs) {
			if (i == len(str) || str[i] != s0[i]) {
				return s0[:i]
			}
		}
	}
	return s0
}