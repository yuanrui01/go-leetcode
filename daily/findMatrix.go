package daily

// 2610. 转换二维数组
func findMatrix(nums []int) [][]int {
	count := make([]int, len(nums)+1)
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	for true {
		arr := []int{}
		for i := 1; i < len(count); i++ {
			if count[i] > 0 {
				arr = append(arr, i)
				count[i]--
			}
		}
		if len(arr) == 0 {
			break
		} else {
			ans = append(ans, arr)
		}
	}
	return ans
}
