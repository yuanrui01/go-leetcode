package daily

// 80. 删除有序数组中的重复项 II
func removeDuplicates(nums []int) int {
	i1, i2, n := 0, 0, len(nums)

	for i1 < n {
		num := nums[i1]
		cnt := 1
		for i1+1 < n && nums[i1+1] == num {
			cnt++
			i1++
		}
		inc2 := min(2, cnt)
		for i := 0; i <= inc2; i++ {
			nums[i2] = num
			i2++
		}
		i1++
	}
	return i2
}
