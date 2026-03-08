# https://leetcode.cn/problems/coin-change/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        # 0 inf inf inf inf inf inf inf inf inf inf inf
        # 0 1   2   3   4   5   6   7   8   9   10  11
        # 0 1   1   2   2   3   3 ...
        # 凑出i元最少需要dp[i]枚硬币
        dp = [float("inf")] * (amount + 1)
        dp[0] = 0

        for coin in coins:
            for j in range(coin, amount + 1):
                dp[j] = min(dp[j], dp[j - coin] + 1)

            print(dp)

        return -1 if dp[amount] == float("inf") else dp[amount]  # type: ignore


if __name__ == "__main__":
    solution = Solution()
    print(solution.coinChange([1, 2, 5], 11))
