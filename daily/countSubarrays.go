package daily

import "slices"

// 2962. 统计最大元素出现至少 K 次的子数组
func countSubarrays2962(nums []int, k int) int64 {
	n, mx := len(nums), slices.Max(nums)
	left, right, ans, mxCnt := 0, 0, 0, 0
	for right < len(nums) {
		if nums[right] == mx {
			mxCnt++
		}
		for mxCnt >= k {
			ans += n - right
			if nums[left] == mx {
				mxCnt--
			}
			left++
		}
		right++
	}
	return int64(ans)
}
