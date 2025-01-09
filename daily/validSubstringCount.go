package daily

// ValidSubstringCount 3297. 统计重新排列后包含另一个字符串的子字符串数目 I
func ValidSubstringCount(word1 string, word2 string) int64 {
	ans := int64(0)

	len1 := len(word1)
	len2 := len(word2)
	if len1 < len2 {
		return ans
	}

	process := make([]int, 26)
	target := make([]int, 26)
	for i := 0; i < len2; i++ {
		target[int(word2[i]-'a')]++
	}
	for i := 0; i < len2-1; i++ {
		process[int(word1[i]-'a')]++
	}

	left := 0
	right := len2 - 2
	for left < len1 {
		for right+1 < len1 && !check(process, target) {
			right++
			process[int(word1[right]-'a')]++
		}
		if check(process, target) {
			ans += int64(len1 - right)
		}
		process[int(word1[left]-'a')]--
		left++
	}
	return ans
}

func check(process []int, target []int) bool {
	for i := 0; i < 26; i++ {
		if target[i] > process[i] {
			return false
		}
	}
	return true
}
