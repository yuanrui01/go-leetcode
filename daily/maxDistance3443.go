package daily

// 3443. K 次修改后的最大曼哈顿距离
func maxDistance3443(s string, k int) int {
	cN, cS, cW, cE := 0, 0, 0, 0
	ans, absV, absH, ck := 0, 0, 0, 0
	for _, c := range s {
		if c == 'N' {
			cN++
		} else if c == 'S' {
			cS++
		} else if c == 'W' {
			cW++
		} else {
			cE++
		}
		absV = abs(cN - cS)
		absH = abs(cW - cE)
		ck = min(k, min(cN, cS)+min(cW, cE))
		ans = max(ans, absV+absH+2*ck)
	}
	return ans
}

// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }
