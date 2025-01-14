package daily

import (
	"container/heap"
	"sort"
)

// 3066. 超过阈值的最少操作数 II
func minOperations3066(nums []int, k int) (ans int) {
	h := &mh{nums}
	heap.Init(h)
	for h.IntSlice[0] < k {
		x := heap.Pop(h).(int)
		h.IntSlice[0] += x * 2
		heap.Fix(h, 0)
		ans++
	}
	return ans
}

type mh struct{ sort.IntSlice }

func (mh) Push(any) {}
func (h *mh) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
