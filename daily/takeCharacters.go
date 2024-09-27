package daily

// TakeCharacters 2516. 每种字符至少取 K 个
func TakeCharacters(s string, k int) (ans int) {
	tm := [3]int{}
	cm := [3]int{}
	for i := 0; i < len(s); i++ {
		tm[s[i]-'a']++
	}
	for i := 0; i < 3; i++ {
		if tm[i] < k {
			return -1
		}
		tm[i] = tm[i] - k
	}
	mx, left := 0, 0
	for right, c := range s {
		c -= 'a'
		cm[c]++
		for cm[c] > tm[c] {
			cm[s[left]-'a']--
			left++
		}
		mx = max(mx, right-left+1)
	}
	return len(s) - mx
}
