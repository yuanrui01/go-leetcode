package daily

// 3305. 元音辅音字符串计数 I
func countOfSubstrings(word string, k int) int {
	n := len(word)
	s := []byte(word)
	ch := make([][6]int, n+1)
	ac, ec, ic, oc, uc, fc := 0, 0, 0, 0, 0, 0
	for i := 0; i < n; i++ {
		if s[i] == 'a' {
			ac++
		} else if s[i] == 'e' {
			ec++
		} else if s[i] == 'i' {
			ic++
		} else if s[i] == 'o' {
			oc++
		} else if s[i] == 'u' {
			uc++
		} else {
			fc++
		}
		ch[i+1][0] = ac
		ch[i+1][1] = ec
		ch[i+1][2] = ic
		ch[i+1][3] = oc
		ch[i+1][4] = uc
		ch[i+1][5] = fc
	}
	var check func(l int, r int) bool
	check = func(l int, r int) bool {
		for i := 0; i < 5; i++ {
			if ch[r][i]-ch[l][i] == 0 {
				return false
			}
		}
		return ch[r][5]-ch[l][5] == k
	}
	ans := 0
	for w := k + 5; w <= n; w++ {
		for i := 0; i <= n-w; i++ {
			if check(i, i+w) {
				ans++
			}
		}
	}
	return ans
}
