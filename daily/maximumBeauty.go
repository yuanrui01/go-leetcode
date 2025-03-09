package daily

import "sort"

// 2070. 每一个查询的最大美丽值
func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		if items[i][0] == items[j][0] {
			return items[i][1] > items[j][1]
		}
		return items[i][0] < items[j][0]
	})
	mx := 0
	s := [][]int{}
	for i := range items {
		if items[i][1] > mx {
			mx = items[i][1]
		}
		items[i][1] = mx

		if len(s) == 0 || s[len(s)-1][0] != items[i][0] {
			s = append(s, items[i])
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		left, right := 0, len(s)-1
		for left <= right {
			mid := (left + right) / 2
			if q < s[mid][0] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if right == -1 {
			ans[i] = 0
		} else {
			ans[i] = s[right][1]
		}
	}
	return ans
}
