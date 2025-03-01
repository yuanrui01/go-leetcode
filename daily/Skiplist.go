package daily

//
//// 1206. 设计跳表
//type Skiplist struct {
//	count [20001]int
//}
//
//func Constructor() *Skiplist {
//	return &Skiplist{}
//}
//
//func (this *Skiplist) Search(target int) bool {
//	return this.count[target] > 0
//}
//
//func (this *Skiplist) Add(num int) {
//	this.count[num]++
//}
//
//func (this *Skiplist) Erase(num int) bool {
//	if this.count[num] == 0 {
//		return false
//	}
//	this.count[num]--
//	return true
//}
