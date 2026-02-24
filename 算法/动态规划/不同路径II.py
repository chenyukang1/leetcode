# https://leetcode.cn/problems/unique-paths-ii/

from typing import List


class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        rows = len(obstacleGrid)
        cols = len(obstacleGrid[0])
        if obstacleGrid[rows - 1][cols - 1] == 1:
            return 0
        dp = [[0 for _ in range(cols)] for _ in range(rows)]
        dp[0][0] = 1
        for i in range(rows):
            for j in range(cols):
                if obstacleGrid[i][j] == 1:
                    continue
                if i - 1 >= 0 and i - 1 < rows and obstacleGrid[i - 1][j] == 0:
                    dp[i][j] += dp[i - 1][j]
                if j - 1 >= 0 and j - 1 < cols and obstacleGrid[i][j - 1] == 0:
                    dp[i][j] += dp[i][j - 1]

        return dp[rows - 1][cols - 1]
