package daily

// 1656. 设计有序流
type OrderedStream struct {
	n   int
	ptr int
	ss  []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{n: n + 1, ptr: 1, ss: make([]string, n+1)}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.ss[idKey] = value
	if this.ss[this.ptr] == "" {
		return []string{}
	}
	ans := []string{}
	for i := this.ptr; i < this.n; i++ {
		if this.ss[i] != "" {
			ans = append(ans, this.ss[i])
			this.ptr = i + 1
		} else {
			break
		}
	}
	return ans
}
