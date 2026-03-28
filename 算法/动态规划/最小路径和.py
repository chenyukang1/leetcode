# https://leetcode.cn/problems/minimum-path-sum/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        rows = len(grid)
        cols = len(grid[0])
        dp = [[0 for _ in range(cols)] for _ in range(rows)]

        for i in range(rows):
            for j in range(cols):
                if i == 0 and j == 0:
                    dp[i][j] = grid[0][0]
                    continue
                last = -1
                if i - 1 >= 0 and j >= 0 and i - 1 < rows and j < cols:
                    last = dp[i - 1][j]
                if i >= 0 and i < rows and j - 1 >= 0 and j - 1 < cols:
                    if last is -1:
                        last = dp[i][j - 1]
                    else:
                        last = min(last, dp[i][j - 1])
                dp[i][j] = last + grid[i][j]

        return dp[rows - 1][cols - 1]
