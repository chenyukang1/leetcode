package 数组

//请你判断一个9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。
//
//数字1-9在每一行只能出现一次。
//数字1-9在每一列只能出现一次。
//数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）
//
//
//注意：
//
//一个有效的数独（部分已被填充）不一定是可解的。
//只需要根据以上规则，验证已经填入的数字是否有效即可。
//空白格用'.'表示。
//

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		visitedRow := make([]bool, 9)
		visitedCol := make([]bool, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				idx := board[i][j] - '0' - 1
				if visitedRow[idx] {
					return false
				}
				visitedRow[idx] = true
			}
			if board[j][i] != '.' {
				idx := board[j][i] - '0' - 1
				if visitedCol[idx] {
					return false
				}
				visitedCol[idx] = true
			}
		}
	}

	for i := 0; i < 9; {
		for j := 0; j < 9; j++ {
			if (j+1)%3 == 0 {
				visited := make([]bool, 9)
				for m := i; m <= m+2; m++ {
					for k := j - 2; k <= j; k++ {
						if board[m][k] != '.' {
							idx := board[m][k] - '0' - 1
							if visited[idx] {
								return false
							}
							visited[idx] = true
						}
					}
				}
			}
		}
		i += 3
	}

	return true
}
