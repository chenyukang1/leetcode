from typing import List


class Solution:
    def partition(self, s: str) -> List[List[str]]:
        ans = []
        path = []
        self.dsf(1, s, path, ans)
        if self.isHuiwen(s):
            ans.append([s])
        return ans

    def dsf(self, start, s, path: List[int], ans):
        if len(path) > 0:
            letters = []
            cur = 0
            for i in path:
                letters.append(s[cur:i])
                cur = i
            letters.append(s[cur:])
            if all(self.isHuiwen(letter) for letter in letters):
                ans.append(letters)

        for i in range(start, len(s)):
            path.append(i)
            self.dsf(i + 1, s, path, ans)
            path.pop()

    def isHuiwen(self, letter):
        i, j = 0, len(letter) - 1
        while i < j:
            if letter[i] != letter[j]:
                return False
            i += 1
            j -= 1
        return True

if __name__ == "__main__":
    solution = Solution()
    result = solution.partition("cdd")
    print((result))
