from typing import List


class Solution:
    def candy(self, ratings: List[int]) -> int:
        arr = [1] * len(ratings)
        # 左规则
        for i in range(1, len(ratings)):
            if ratings[i - 1] < ratings[i]:
                arr[i] = arr[i - 1] + 1
            else:
                arr[i] = 1

        # 右规则
        for i in range(len(ratings) - 2, -1, -1):
            if ratings[i] > ratings[i + 1]:
                arr[i] = max(arr[i], arr[i + 1] + 1)
            else:
                arr[i + 1] = max(arr[i + 1], 1)

        return sum(arr)
