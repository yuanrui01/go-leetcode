package daily

// 9. 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arr := [12]int{}
	l := 0
	for x != 0 {
		arr[l] = x % 10
		x /= 10
		l++
	}
	mid := l / 2
	for i := 0; i < mid; i++ {
		if arr[i] != arr[l-i-1] {
			return false
		}
	}
	return true
}
