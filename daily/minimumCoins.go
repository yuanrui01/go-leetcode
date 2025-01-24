package daily

import "math"

// 2944. 购买水果需要的最少金币数
func minimumCoins(prices []int) int {
	n := len(prices)
	memo := make([]int, (n+1)/2)
	return dp(1, prices, memo)
}

func dp(i int, prices []int, memo []int) int {
	if 2*i >= len(prices) {
		return prices[i-1]
	}
	if memo[i] != 0 {
		return memo[i]
	}
	res := math.MaxInt32
	for j := i + 1; j <= 2*i+1; j++ {
		res = min(res, dp(j, prices, memo))
	}
	memo[i] = res + prices[i-1]
	return memo[i]
}
