package daily

// NumArray 303. 区域和检索 - 数组不可变
type NumArray []int

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func (a NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return a[right]
	}
	return a[right] - a[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
