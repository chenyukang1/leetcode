# https://leetcode.cn/problems/ones-and-zeroes/

from typing import List


class Solution:
    def findMaxForm(self, strs: List[str], m: int, n: int) -> int:
        # dp[i][j] 表示容量i个0，j个1最大子集的长度
        # dp[0][0] = 0
        # dp[i][j] = max(dp[i][j], dp[i-len_0(strs[i])][j-len_1(strs[i])]+1)
        dp = [[0 for _ in range(n + 1)] for _ in range(m + 1)]
        for i in range(len(strs)):
            nums_0 = strs[i].count("0")
            nums_1 = strs[i].count("1")
            for j in range(m, nums_0 - 1, -1):
                for k in range(n, nums_1 - 1, -1):
                    dp[j][k] = max(dp[j][k], dp[j - nums_0][k - nums_1] + 1)

        return dp[m][n]
