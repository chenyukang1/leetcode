package _62

// Q1. 垂直翻转子矩阵
// 给你一个 m x n 的整数矩阵 grid，以及三个整数 x、y 和 k。
//
// 整数 x 和 y 表示一个 正方形子矩阵 的左上角下标，整数 k 表示该正方形子矩阵的边长。
//
// 你的任务是垂直翻转子矩阵的行顺序。
//
// 返回更新后的矩阵。©leetcode
func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	i := x
	j := x + k - 1
	for i < j {
		temp := make([]int, k)
		for m, n := y, 0; m < y+k; m, n = m+1, n+1 {
			temp[n] = grid[i][m]
			grid[i][m] = grid[j][m]
			grid[j][m] = temp[n]
		}
		i++
		j--
	}
	return grid
}
