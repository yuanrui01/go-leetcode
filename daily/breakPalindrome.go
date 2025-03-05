package daily

// 1328. 破坏回文串
func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	if n == 1 {
		return ""
	}
	s := []byte(palindrome)
	isOdd := n%2 == 1
	mid := n / 2
	for i := 0; i < n; i++ {
		if isOdd && i == mid {
			continue
		}
		if s[i] != 'a' {
			s[i] = 'a'
			return string(s)
		}
	}
	s[n-1]++
	return string(s)
}
