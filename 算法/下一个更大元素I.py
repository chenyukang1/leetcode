# https://leetcode.cn/problems/next-greater-element-i/

from typing import List


class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        dict = {}
        for i, num in enumerate(nums2):
            dict[num] = i

        next = {}
        next[nums2[-1]] = -1

        stack = [nums2[-1]]
        for i in range(len(nums2) - 2, -1, -1):
            while stack and stack[-1] < nums2[i]:
                stack.pop()
            if stack:
                next[nums2[i]] = stack[-1]
            else:
                next[nums2[i]] = -1
            stack.append(nums2[i])

        print(next)

        return [next[i] for i in nums1]


if __name__ == "__main__":
    s = Solution().nextGreaterElement([4, 1, 2], [1, 3, 4, 2])
    print(s)
