package hot100

import "slices"

// https://leetcode.cn/problems/merge-intervals/?envType=study-plan-v2&envId=top-100-liked

func merge(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	for _, interval := range intervals {
		l := len(ans)
		if l > 0 && ans[l-1][1] >= interval[0] {
			ans[l-1][1] = max(ans[l-1][1], interval[1])
		} else {
			ans = append(ans, interval)
		}
	}
	return
}
