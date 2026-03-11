# https://leetcode.cn/problems/perfect-squares/?envType=study-plan-v2&envId=top-100-liked


class Solution:
    def numSquares(self, n: int) -> int:
        # 和为i的完全平方数最少数量为dp[i]
        # nums = [1,4,9]
        # 0 0 0 0 0 0 0 0 0 0 0  0  0
        # 0 1 2 3 4 5 6 7 8 9 10 11 12
        # 0 1 2 3 1 2 3 4 2 3 4  5  3
        # 0 1 2 3 1 2 3 4 2 1 2  3  3
        dp = [float("inf")] * (n + 1)
        dp[0] = 0

        nums = []
        idx, num = 1, 1
        while num <= n:
            nums.append(num)
            idx += 1
            num = idx**2

        for num in nums:
            for j in range(num, n + 1):
                dp[j] = min(dp[j], dp[j - num] + 1)

        return dp[n]  # type: ignore


if __name__ == "__main__":
    solution = Solution()
    print(solution.numSquares(12))
