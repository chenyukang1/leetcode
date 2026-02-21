# https://leetcode.cn/problems/pascals-triangle/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        ans = []
        for i in range(numRows):
            arr = [1] * (i + 1)
            for j in range(1, i):
                arr[j] = ans[i - 1][j - 1] + ans[i - 1][j]
            ans.append(arr)

        return ans
