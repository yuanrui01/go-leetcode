package daily

import "slices"

// 3170. 删除星号以后字典序最小的字符串
func clearStars(s string) string {
	stacks := make([][]int, 26)
	for i, c := range s {
		if c != '*' {
			stacks[c-'a'] = append(stacks[c-'a'], i)
			continue
		}
		for j, st := range stacks {
			if len(st) > 0 {
				stacks[j] = st[:len(st)-1]
				break
			}
		}
	}

	idx := []int{}
	for _, p := range stacks {
		idx = append(idx, p...)
	}
	slices.Sort(idx)
	res := make([]byte, len(idx))
	for i, v := range idx {
		res[i] = s[v]
	}
	return string(res)
}
