package daily

// 2412. 完成所有交易的初始最少钱数
func minimumMoney(transactions [][]int) int64 {
	ans := int64(0)
	mx := 0
	for _, tx := range transactions {
		ans += int64(max(0, tx[0]-tx[1]))
		mx = max(mx, min(tx[0], tx[1]))
	}
	return ans + int64(mx)
}
