package daily

import (
	"container/heap"
	"sort"
)

// SeatManager 1845. 座位预约管理系统
type SeatManager struct {
	sort.IntSlice
}

func Constructor(n int) SeatManager {
	m := SeatManager{make([]int, n)}
	for i := range m.IntSlice {
		m.IntSlice[i] = i + 1
	}
	return m
}

func (m *SeatManager) Reserve() int {
	return heap.Pop(m).(int)
}

func (m *SeatManager) Unreserve(seatNumber int) {
	heap.Push(m, seatNumber)
}

func (m *SeatManager) Push(v any) { m.IntSlice = append(m.IntSlice, v.(int)) }
func (m *SeatManager) Pop() any {
	a := m.IntSlice
	v := a[len(a)-1]
	m.IntSlice = a[:len(a)-1]
	return v
}
