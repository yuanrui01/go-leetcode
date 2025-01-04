package mycalendarthree

import "sort"

// MyCalendarThree 732. 我的日程安排表 III
type MyCalendarThree struct {
	count map[int]int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		count: make(map[int]int),
	}
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	this.count[startTime]++
	this.count[endTime]--
	ans := 0
	sum := 0

	keys := []int{}
	for k := range this.count {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for k := range keys {
		sum += this.count[keys[k]]
		ans = max(ans, sum)
	}
	return ans
}
