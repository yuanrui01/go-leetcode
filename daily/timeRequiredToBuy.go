package daily

// 2073. 买票需要的时间
func timeRequiredToBuy(tickets []int, k int) int {
	ans := 0
	tk := tickets[k]
	for i, t := range tickets {
		if i > k {
			ans += min(t, tk-1)
		} else {
			ans += min(t, tk)
		}
	}
	return ans
}
