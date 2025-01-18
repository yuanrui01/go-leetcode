package daily


// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
    ans := 0
    right, left := 0, 0
    mark := make(map[uint8]int)
    for right < len(s) {
    	value, exists := mark[s[right]]
    	if exists {
    		left = max(left, value + 1)
    	}
    	mark[s[right]] = right
    	ans = max(ans, right-left+1)
    	right++
    }
    return ans
}