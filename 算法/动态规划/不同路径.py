# https://leetcode.cn/problems/unique-paths/


class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        dp = [[0 for _ in range(n)] for _ in range(m)]
        dp[0][0] = 1
        for i in range(m):
            for j in range(n):
                if j - 1 >= 0 and j - 1 < n:
                    dp[i][j] += dp[i][j - 1]
                if i - 1 >= 0 and i - 1 < m:
                    dp[i][j] += dp[i - 1][j]

        return dp[m - 1][n - 1]
