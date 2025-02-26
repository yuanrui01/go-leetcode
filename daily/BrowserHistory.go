package daily

// BrowserHistory 1472. 设计浏览器历史记录
type BrowserHistory struct {
	his  []string
	cur  int
	last int
}

//func Constructor(homepage string) BrowserHistory {
//	b := BrowserHistory{make([]string, 5001), 0, 0}
//	b.his[0] = homepage
//	return b
//}

func (this *BrowserHistory) Visit(url string) {
	this.cur++
	this.his[this.cur] = url
	this.last = this.cur
}

func (this *BrowserHistory) Back(steps int) string {
	pos := max(0, this.cur-steps)
	page := this.his[pos]
	this.cur = pos
	return page
}

func (this *BrowserHistory) Forward(steps int) string {
	pos := min(this.last, this.cur+steps)
	page := this.his[pos]
	this.cur = pos
	return page
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
