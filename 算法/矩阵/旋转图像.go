package 矩阵

// https://leetcode.cn/problems/rotate-image/?envType=study-plan-v2&envId=top-100-liked

func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l/2; i++ {
		for j := i; j < l-i-1; j++ {
			matrix[i][j], matrix[l-j-1][i], matrix[l-i-1][l-j-1], matrix[j][l-i-1] =
				matrix[l-j-1][i], matrix[l-i-1][l-j-1], matrix[j][l-i-1], matrix[i][j]
		}
	}
}
