package daily

// 2080. 区间内查询数字的频率
//type RangeFreqQuery map[int][]int
//
//func Constructor(arr []int) RangeFreqQuery {
//	pos := map[int][]int{}
//	for i, x := range arr {
//		pos[x] = append(pos[x], i)
//	}
//	return pos
//}
//
//func (pos RangeFreqQuery) Query(left int, right int, value int) int {
//	a := pos[value]
//	return sort.SearchInts(a, right+1) - sort.SearchInts(a, left)
//}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
