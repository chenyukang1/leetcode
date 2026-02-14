package 矩阵

// https://leetcode.cn/problems/search-a-2d-matrix-ii/?envType=study-plan-v2&envId=top-100-liked

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])
	l := row
	if row > col {
		l = col
	}
	for i := 0; i < l; i++ {
		left := i
		right := col - i - 1
		// [left, right]
		for left <= right {
			mid := left + (right-left)>>1
			if matrix[i][mid] < target {
				left = mid + 1
			} else if matrix[i][mid] > target {
				right = mid - 1
			} else {
				return true
			}
		}

		left = i
		right = row - i - 1
		// [left, right]
		for left <= right {
			mid := left + (right-left)>>1
			if matrix[mid][i] < target {
				left = mid + 1
			} else if matrix[mid][i] > target {
				right = mid - 1
			} else {
				return true
			}
		}
	}

	return false
}

// 从右上角开始遍历 O(m+n)
func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}
