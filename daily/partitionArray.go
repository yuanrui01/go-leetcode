package daily

func PartitionArray(nums []int, k int) (ans int) {
	index, rem, n := 0, k, len(nums)

	for index < n {
		for index+1 < n && nums[index+1]-nums[index] <= rem {
			rem -= nums[index+1] - nums[index]
			index++
		}
		rem = k
		ans++
		index++
	}
	return
}
