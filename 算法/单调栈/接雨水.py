# https://leetcode.cn/problems/trapping-rain-water/description/?envType=study-plan-v2&envId=top-100-liked

import collections
from typing import List


class Solution:
    # 思路一: 单调栈，维护一个递减的单调栈，遇到比栈顶高的元素出栈，取现在栈顶和当前元素的较小值为可以蓄水的长度
    def trap(self, height: List[int]) -> int:
        ans = 0
        stack = collections.deque()
        stack.append(0)
        for i in range(1, len(height)):
            while stack and height[stack[-1]] < height[i]:
                cur_idx = stack.pop()
                if stack:
                    short = min(height[stack[-1]], height[i])
                    area = (short - height[cur_idx]) * (i - stack[-1] - 1)
                    ans += area
            stack.append(i)

        return ans

    # 思路2: DP，两次遍历取从左/右看能取到的最大高度。最后遍历每个量柱，取左/右看的较小者和当前高度的差值:w
    def trap2(self, height: List[int]) -> int:
        n = len(height)
        left_max = [height[0]] + [0] * (n - 1)
        right_max = [0] * (n - 1) + [height[n - 1]]

        for i in range(1, n):
            left_max[i] = max(left_max[i - 1], height[i])

        for i in range(n - 2, -1, -1):
            right_max[i] = max(right_max[i + 1], height[i])

        ans = 0
        for i in range(n):
            ans += min(left_max[i], right_max[i]) - height[i]

        return ans


if __name__ == "__main__":
    solution = Solution()
    print(solution.trap2([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]))
