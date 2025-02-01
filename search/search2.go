package search

// 81. 搜索旋转排序数组 II
func search2(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	last := nums[right]
	for left <= right {
		mid := left + (right-left)/2
		nm := nums[mid]
		if nm == target {
			return true
		}
		if nm == nums[left] {
			left++
		} else if nm > last {
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

	return nums[right] == target
}
