# https://leetcode.cn/problems/last-stone-weight-ii/

from typing import List


class Solution:
    def lastStoneWeightII(self, stones: List[int]) -> int:
        all_sum = sum(stones)
        target = all_sum // 2
        dp = [0] * (target + 1)
        # dp[i] 表示放入总重量i的背包，总共石头的最大重量
        # dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
        for i in range(len(stones)):
            for j in range(target, stones[i] - 1, -1):
                dp[j] = max(dp[j], dp[j - stones[i]] + stones[i])

        return all_sum - dp[target] * 2
