package hot100

// https://leetcode.cn/problems/set-matrix-zeroes/description/?envType=study-plan-v2&envId=top-100-liked

// setZeroes 时间复杂度O(m*n) 空间复杂度O(m+n)
func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	rows := make([]bool, m)
	cols := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i, r := range matrix {
		for j := range r {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// setZeroes 时间复杂度O(m*n) 空间复杂度O(1)
func setZeroes2(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	row0, col0 := false, false
	for _, i := range matrix[0] {
		if i == 0 {
			row0 = true
			break
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = true
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if col0 {
		for _, r := range matrix {
			r[0] = 0
		}
	}
}
