package daily

// 2799. 统计完全子数组的数目
func countCompleteSubarrays(nums []int) (ans int) {
	set := map[int]struct{}{}
	for _, n := range nums {
		set[n] = struct{}{}
	}

	n := len(nums)
	k := len(set)

	cnt := make(map[int]int, k)
	left, right := 0, 0
	for right < n {
		cnt[nums[right]]++
		if len(cnt) == k {
			inc := n - right
			for left <= right && len(cnt) == k {
				ans += inc
				cnt[nums[left]]--
				if cnt[nums[left]] == 0 {
					delete(cnt, nums[left])
				}
				left++
			}
		}
		right++
	}
	return ans
}
