package daily

// 680. Valid Palindrome II
func validPalindrome(s string) bool {
	var valid func(s string, left int, right int) bool
	valid = func(s string, left int, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return valid(s, left+1, right) || valid(s, left, right-1)
		}
		left++
		right--
	}
	return true
}
