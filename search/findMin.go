package search

// 153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	last := nums[right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= last {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
