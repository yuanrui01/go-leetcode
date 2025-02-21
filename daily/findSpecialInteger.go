package daily

// 1287. 有序数组中出现次数超过25%的元素
func findSpecialInteger(arr []int) int {
	n := len(arr)
	limit := n / 4
	for i := 0; i < n; i++ {
		cnt := 1
		num := arr[i]
		for i+1 < n && arr[i+1] == num {
			cnt++
			i++
		}
		if cnt > limit {
			return num
		}
	}
	return -1
}
