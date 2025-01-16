package daily

import "math"

// MinimumSubarrayLength 3097. 或值至少为 K 的最短子数组 II
func MinimumSubarrayLength(nums []int, k int) int {
	left, right := 0, 0
	bitCount := make([]int, 31)
	length := len(nums)
	ans := math.MaxInt32
	for right < length {
		addBits(bitCount, nums[right])
		for left <= right && getDNum(bitCount) >= k {
			ans = min(ans, right-left+1)
			removeBits(bitCount, nums[left])
			left++
		}
		right++
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func addBits(bitCount []int, num int) {
	for i := 0; i < 31; i++ {
		bitCount[i] += num & 1
		num >>= 1
	}
}

func removeBits(bitCount []int, num int) {
	for i := 0; i < 31; i++ {
		bitCount[i] -= num & 1
		num >>= 1
	}
}

func getDNum(bitCount []int) int {
	factor := 1
	num := 0
	for i := 0; i < 31; i++ {
		if bitCount[i] != 0 {
			num += factor
		}
		factor <<= 1
	}
	return num
}
