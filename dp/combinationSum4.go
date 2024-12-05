package dp

// 377. 组合总和 Ⅳ
func combinationSum4(nums []int, target int) int {
	cache := make([]int, target+1)
	for i := 0; i < target+1; i++ {
		cache[i] = -1
	}
	return dfs(nums, cache, target)
}

func dfs(nums []int, cache []int, target int) int {
	if target == 0 {
		return 1
	}
	if cache[target] != -1 {
		return cache[target]
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		if target-nums[i] >= 0 {
			ans += dfs(nums, cache, target-nums[i])
		}
	}
	cache[target] = ans
	return ans
}
