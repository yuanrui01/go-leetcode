package daily

func lengthOfLongestSubstring(s string) int {
	ans := 0
	left, right := 0, 0
	count := make(map[uint8]int)
	for right < len(s) {
		value, exists := count[s[right]]
		if exists {
			left = max(left, value+1)
		}
		count[s[right]] = right
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}
