from typing import List


class Solution:
    # 思路一：记忆化回溯
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        word_set = set(wordDict)
        memo = {}

        def backTrack(start):
            if start == len(wordDict):
                return True
            if start in memo:
                return False

            for i in range(start + 1, len(s) + 1):
                word = s[start:i]
                if word in word_set and backTrack(i):
                    return True

            memo[start] = False

            return False

        return backTrack(0)

    def wordBreak2(self, s: str, wordDict: List[str]) -> bool:
        # dp[i]表示到i为止能否拆成子单词
        # dp[i] = dp[i-j] and s[i-j+1:i] in wordDict
        word_set = set(wordDict)
        dp = [False] * (len(s) + 1)
        dp[0] = True

        for i in range(1, len(s) + 1):
            for j in range(1, i + 1):
                if dp[i - j] and s[i - j : i] in word_set:
                    dp[i] = True
                    continue

        return dp[len(s)]


if __name__ == "__main__":
    s = Solution()
    res = s.wordBreak2("leetcode", ["leet", "code"])
    print(res)
