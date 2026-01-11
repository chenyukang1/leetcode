package 排序和搜索

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	l, r := 1, n
	// 闭区间[l,r]
	for l < r {
		mid := l + (r-l)>>1
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
