package daily

// 2597. 美丽子集的数目
func beautifulSubsets(nums []int, k int) int {
	cnt := map[int]int{}
	n := len(nums)
	ans := -1
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans++
			return
		}
		dfs(i + 1)
		x := nums[i]
		if cnt[x-k] == 0 && cnt[x+k] == 0 {
			cnt[x]++
			dfs(i + 1)
			cnt[x]--
		}
	}
	dfs(0)
	return ans
}
