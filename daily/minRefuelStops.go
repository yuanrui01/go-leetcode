package daily

import (
	"container/heap"
	"sort"
)

// 871. 最低加油次数
func minRefuelStops(target int, startFuel int, stations [][]int) (ans int) {
	stations = append(stations, []int{target, 0})
	prePosition, curFuel := 0, startFuel
	fuelMap := &hp{}
	for _, station := range stations {
		position, fuel := station[0], station[1]
		curFuel -= position - prePosition
		for fuelMap.Len() > 0 && curFuel < 0 {
			curFuel += heap.Pop(fuelMap).(int)
			ans++
		}
		if curFuel < 0 {
			return -1
		}
		heap.Push(fuelMap, fuel)
		prePosition = position
	}
	return ans
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
