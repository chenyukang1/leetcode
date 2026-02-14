# https://leetcode.cn/problems/0ynMMM/description/

from typing import List


class Solution:
    # 遍历每个柱子，向左右延伸最大距离
    # O(n^2) 超时，不能通过
    def largestRectangleArea(self, heights: List[int]) -> int:
        max_area = 0
        n = len(heights)
        for i in range(n):
            left, right = i, i
            while left-1 >= 0 and heights[left-1] >= heights[i]:
                left -= 1
            while right+1 < n and heights[right+1] >= heights[i]:
                right += 1
            max_area = max(max_area, (right - left + 1) * heights[i])

        return max_area

    # 单调栈，每次和栈顶元素比较计算最大面积
    # O(n)
    def largestRectangleArea_v2(self, heights: List[int]) -> int:
        if not heights:
            return 0
        max_area = 0
        n = len(heights)
        heights.append(-1)
        stack = [0]
        for i in range(1, n + 1):
            while len(stack) > 0 and heights[i] < heights[stack[-1]]:
                j = stack.pop()
                left = -1 if len(stack) == 0 else stack[-1]
                max_area = max(max_area, (i-left-1) * heights[j])
            stack.append(i)

        return max_area


if __name__ == "__main__":
    solution = Solution()
    print(solution.largestRectangleArea([2,1,5,6,2,3]))
    print(solution.largestRectangleArea_v2([2,1,5,6,2,3]))
