package daily

// 1436. 旅行终点站
func destCity(paths [][]string) string {
	setA := make(map[string]struct{}, len(paths))
	for _, p := range paths {
		setA[p[0]] = struct{}{}
	}

	for _, p := range paths {
		if _, ok := setA[p[1]]; !ok {
			return p[1]
		}
	}
	return ""
}
