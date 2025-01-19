package daily

const mod = 1_000_000_007

var cache3 = make(map[int]int)
var cache4 = make(map[int]int)
var cm = make(map[uint8]int)

func init() {
	cm['2'] = 3
	cm['3'] = 3
	cm['4'] = 3
	cm['5'] = 3
	cm['6'] = 3
	cm['7'] = 4
	cm['8'] = 3
	cm['9'] = 4
}

// CountTexts 2266. 统计打字方案数
func CountTexts(pressedKeys string) int {
	length := len(pressedKeys)
	i := 0
	result := 1
	for i < length {
		c := pressedKeys[i]
		conc := 1
		for i+1 < length && pressedKeys[i+1] == c {
			conc++
			i++
		}
		if cm[c] == 3 {
			result = (result * dp3(conc)) % mod
		} else {
			result = (result * dp4(conc)) % mod
		}
		i++
	}
	return result
}

func dp3(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	value, exists := cache3[n]
	if exists {
		return value
	}
	result := 0
	for i := 1; i <= 3; i++ {
		if n-i >= 0 {
			result = (result + dp3(n-i)) % mod
		}
	}
	cache3[n] = result
	return result
}

func dp4(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	value, exists := cache4[n]
	if exists {
		return value
	}
	result := 0
	for i := 1; i <= 4; i++ {
		if n-i >= 0 {
			result = (result + dp4(n-i)) % mod
		}
	}
	cache4[n] = result
	return result
}
