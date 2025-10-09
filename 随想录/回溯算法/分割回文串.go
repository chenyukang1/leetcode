package 回溯算法

func partition(s string) [][]string {
	var ans [][]string
	var split []string
	backTracking7(0, s, split, &ans)
	return ans
}

func backTracking7(start int, s string, split []string, ans *[][]string) {
	if start == len(s) {
		temp := make([]string, len(split))
		copy(temp, split)
		*ans = append(*ans, temp)
		return
	}

	for i := start + 1; i <= len(s); i++ {
		substr := s[start:i]
		if isHuiWen(substr) {
			split = append(split, substr)
			backTracking7(i, s, split, ans)
			split = split[:len(split)-1]
		}
	}
}

func isHuiWen(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
