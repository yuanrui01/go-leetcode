package daily

import "math"

// MinimumDifference 3171. 找到按位或最接近 K 的子数组
func MinimumDifference(nums []int, k int) int {
	left, right := 0, 0
	n := len(nums)
	digits := make([]int, 32)
	ans := math.MaxInt32
	for right < n {
		alterDigits(digits, nums[right], 1)
		ans = min(ans, abs(k-getNum(digits)))
		for left < right && getNum(digits) > k {
			alterDigits(digits, nums[left], -1)
			ans = min(ans, abs(k-getNum(digits)))
			left++
		}
		right++
	}
	return ans
}

func getNum(digits []int) int {
	sum := 0
	factor := 1
	for _, v := range digits {
		if v > 0 {
			sum += factor
		}
		factor <<= 1
	}
	return sum
}

func alterDigits(digits []int, n int, flag int) {
	i := 0
	for n != 0 {
		digits[i] += flag * (n % 2)
		i++
		n >>= 1
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
