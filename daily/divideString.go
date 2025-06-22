package daily


// 2138. 将字符串拆分为若干长度为 k 的组
func divideString(s string, k int, fill byte) []string {
	n := len(s)
	ans := make([]string, 0, (n+k-1)/k)
	for i := 0; i < n; i += k {
		if i + k <= n {
			ans = append(ans, s[i:i+k])
		} else {
			ans = append(ans, s[i:] + strings.Repeat(string(fill), k-n+i))
		}
	}
	return ans
}