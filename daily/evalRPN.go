package daily

import "strconv"

// LCR 036. 逆波兰表达式求值
func evalRPN(tokens []string) int {
	st := make([]int, 0)
	for _, t := range tokens {
		if t == "+" || t == "-" || t == "*" || t == "/" {
			p1 := st[len(st)-1]
			p2 := st[len(st)-2]
			st = st[:len(st)-2]
			if t == "+" {
				st = append(st, p2+p1)
			} else if t == "-" {
				st = append(st, p2-p1)
			} else if t == "*" {
				st = append(st, p2*p1)
			} else {
				st = append(st, p2/p1)
			}
		} else {
			num, _ := strconv.Atoi(t)
			st = append(st, num)
		}
	}
	return st[0]
}
