package daily

// 1299. 将每个元素替换为右侧最大元素
func replaceElements(arr []int) []int {
	mx := -1
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		tmp := arr[i]
		arr[i] = mx
		mx = max(mx, tmp)
	}
	return arr
}
