package daily

// 1920. 基于排列构建数组
func buildArray(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		if x < 0 {
			continue
		}
		cur := i
		for nums[cur] != i {
			nxt := nums[cur]
			nums[cur] = ^nums[nxt]
			cur = nxt
		}
		nums[cur] = ^x
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = ^nums[i]
	}
	return nums
}

func main() {

}
