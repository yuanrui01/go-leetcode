package daily

import (
	"container/heap"
	"math"
)

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 3341. 到达最后一个房间的最少时间 I
func minTimeToReach(moveTime [][]int) int {
	n, m := len(moveTime), len(moveTime[0])
	t := make([][]int, n)
	for i := range t {
		t[i] = make([]int, m)
		for j := range t[i] {
			t[i][j] = math.MaxInt
		}
	}
	t[0][0] = 0
	h := hp{{}}
	for {
		top := heap.Pop(&h).(tuple)
		i, j := top.x, top.y
		if i == n-1 && j == m-1 {
			return top.t
		}
		if top.t > t[i][j] {
			continue
		}
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if x >= 0 && x < n && y >= 0 && y < m {
				newT := max(top.t, moveTime[x][y]) + 1
				if newT < t[i][j] {
					t[i][j] = newT
					heap.Push(&h, tuple{x, y, newT})
				}
			}
		}
	}
}

type tuple struct{ x, y, t int }
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].t < h[j].t }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() (v any)      { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
