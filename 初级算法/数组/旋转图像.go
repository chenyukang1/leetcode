package 数组

//给定一个 n×n 的二维矩阵matrix 表示一个图像。请你将图像顺时针旋转 90 度
//
//你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像
//
//提示：
//
//n == matrix.length == matrix[i].length
//1 <= n <= 20
//-1000 <= matrix[i][j] <= 1000

// 1 2 3
// 4 5 6
// 7 8 9

// 1 4 7
// 2 5 8
// 3 6 9
// 先做一次斜对称，然后reserve
func rotate2(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}

	for i := 0; i < l; i++ {
		for left, right := 0, l-1; left < right; left, right = left+1, right-1 {
			temp := matrix[i][left]
			matrix[i][left] = matrix[i][right]
			matrix[i][right] = temp
		}
	}
}

// 模拟法，取左上角变量，依次逆时针赋值
func rotate3(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < (l >> 1); i++ {
		for j := i; j < l-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[l-j-1][i]
			matrix[l-j-1][i] = matrix[l-i-1][l-j-1]
			matrix[l-i-1][l-j-1] = matrix[j][l-i-1]
			matrix[j][l-i-1] = temp
		}
	}
}
