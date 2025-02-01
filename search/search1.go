package search

// 33. 搜索旋转排序数组
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	last := nums[right]
	for left <= right {
		mid := left + (right-left)/2
		nm := nums[mid]
		if nm == target {
			return mid
		}
		if nm > last {
			if target > last && nm > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > last || nm > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}
