# https://leetcode.cn/problems/daily-temperatures/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        if len(temperatures) == 1:
            return [0]
        stack = [(len(temperatures) - 1, temperatures[-1])]
        ans = [0] * len(temperatures)
        for i in range(len(temperatures) - 2, -1, -1):
            if not stack:
                stack.append((i, temperatures[i]))
            else:
                while stack[-1][1] < temperatures[i]:
                    stack.pop()
                if stack:
                    ans[i] = stack[-1][0] - i
                stack.append((i, temperatures[i]))

        return ans


if __name__ == "__main__":
    solution = Solution()
    print(solution.dailyTemperatures([73, 74, 75, 71, 69, 72, 76, 73]))
