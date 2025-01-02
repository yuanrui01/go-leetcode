package main

import "fmt"

type MyCalendar struct {
	booked [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{
		booked: [][]int{},
	}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	for _, book := range this.booked {
		if startTime < book[1] && endTime > book[0] {
			return false
		}
	}
	this.booked = append(this.booked, []int{startTime, endTime})
	return true
}

func main() {
	// 创建 MyCalendar 实例
	calendar := Constructor()

	// 测试 Book 方法
	fmt.Println(calendar.Book(10, 20)) // true
	fmt.Println(calendar.Book(15, 25)) // false
	fmt.Println(calendar.Book(20, 30)) // true
}
