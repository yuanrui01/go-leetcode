package daily

// CountKeyChanges 3019. 按键变更的次数
func CountKeyChanges(s string) int {
	ans := 0
	for i := 1; i < len(s); i++ {
		diff := abs(int(s[i]) - int(s[i-1]))
		if diff != 0 && diff != 32 {
			ans++
		}
	}
	return ans
}
