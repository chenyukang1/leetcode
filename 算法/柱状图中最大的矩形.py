# https://leetcode.cn/problems/largest-rectangle-in-histogram/description/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        stack = [0]
        max_area = heights[0]
        for i in range(1, len(heights)):
            max_area = max(max_area, heights[i])
            while stack and heights[stack[-1]] > heights[i]:
                cur_idx = stack.pop()
                left = stack[-1] if stack else -1
                max_area = max(max_area, (i - left - 1) * heights[cur_idx])
            stack.append(i)

        right = len(heights)
        while stack:
            cur_idx = stack.pop()
            left = stack[-1] if stack else -1
            max_area = max(max_area, (right - left - 1) * heights[cur_idx])

        return max_area


if __name__ == "__main__":
    solution = Solution()
    res = solution.largestRectangleArea([2, 1, 5, 6, 2, 3])
    print(res)
