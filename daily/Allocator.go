package daily

type Allocator []int

// 2502. 设计内存分配器
//func Constructor(n int) Allocator {
//	return make([]int, n)
//}

func (a Allocator) Allocate(size int, mID int) int {
	free := 0
	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			free = 0
			continue
		}
		free++
		if free == size {
			for j := i - size + 1; j <= i; j++ {
				a[j] = mID
			}
			return i - size + 1
		}
	}
	return -1
}

func (a Allocator) FreeMemory(mID int) int {
	ans := 0
	for i := 0; i < len(a); i++ {
		if a[i] == mID {
			a[i] = 0
			ans++
		}
	}
	return ans
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.FreeMemory(mID);
 */
