package daily

import "sort"

// 1561. 你可以获得的最大硬币数目
func maxCoins(piles []int) int {
	ans := 0
	sort.Ints(piles)
	pl := len(piles)
	for i := pl / 3; i < pl; i += 2 {
		ans += piles[i]
	}
	return ans
}
