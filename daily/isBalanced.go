package daily

// 3340. 检查平衡字符串
func isBalanced(num string) bool {
	s := []byte(num)
	sum := 0
	for i, b := range s {
		if i%2 == 0 {
			sum += int(b - '0')
		} else {
			sum -= int(b - '0')
		}
	}
	return sum == 0
}
