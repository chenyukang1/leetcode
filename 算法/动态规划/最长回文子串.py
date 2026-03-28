# https://leetcode.cn/problems/longest-palindromic-substring/?envType=study-plan-v2&envId=top-100-liked


class Solution:
    def longestPalindrome(self, s: str) -> str:
        # dp[i][j] 表示从i..j是否是回文串
        dp = [[False for _ in range(len(s))] for _ in range(len(s))]
        for i in range(len(s)):
            dp[i][i] = True

        ans = s[0]
        for sub_len in range(2, len(s) + 1):
            for left in range(len(s)):
                right = left + sub_len - 1
                if right > len(s) - 1:
                    continue

                if s[left] == s[right]:
                    if right - left <= 2:
                        dp[left][right] = True
                    else:
                        dp[left][right] = dp[left + 1][right - 1]

                if dp[left][right] and right - left + 1 > len(ans):
                    ans = s[left : right + 1]

        return ans


if __name__ == "__main__":
    s = Solution()
    res = s.longestPalindrome("aaaa")
    print(res)
