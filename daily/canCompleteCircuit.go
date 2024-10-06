package daily

// 134. 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	ans := 0
	cal := 0
	sMin := 0
	for i, v := range gas {
		cal += v - cost[i]
		if cal < sMin {
			sMin = cal
			ans = i + 1
		}
	}
	if cal < 0 {
		return -1
	} else {
		return ans
	}
}
