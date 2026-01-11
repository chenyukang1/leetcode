package hot100

// https://leetcode.cn/problems/spiral-matrix/?envType=study-plan-v2&envId=top-100-liked

// spiralOrder 层次写法，没法处理中心点
func spiralOrder(matrix [][]int) (ans []int) {
	m := len(matrix)
	n := len(matrix[0])
	layers := (min(m, n) + 1) / 2

	for l := 0; l < layers; l++ {
		for i := l; i < n-l-1; i++ {
			ans = append(ans, matrix[l][i])
		}
		for i := l; i < m-l-1; i++ {
			ans = append(ans, matrix[i][n-l-1])
		}
		if m-l-1 != l {
			for i := n - l - 1; i > l; i-- {
				ans = append(ans, matrix[m-l-1][i])
			}
		}
		if n-l-1 != l {
			for i := m - l - 1; i > l; i-- {
				ans = append(ans, matrix[i][l])
			}
		}
	}
	return ans
}

// spiralOrder2 四边界法
func spiralOrder2(matrix [][]int) (ans []int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, n-1, 0, m-1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--

		if top < bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
		}
		bottom--

		if left < right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
		}
		left++
	}

	return ans
}
