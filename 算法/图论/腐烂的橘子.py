# https://leetcode.cn/problems/rotting-oranges/description/?envType=study-plan-v2&envId=top-100-liked

import collections
from typing import List


class Solution:
    # 思路1: 模拟法，每轮腐烂1次，直到这轮没有新的橘子腐烂，查看最后有没有好橘子
    def orangesRotting(self, grid: List[List[int]]) -> int:
        self.visited = [[False for _ in row] for row in grid]
        ans = 0
        while self.bsf(grid):
            ans += 1

        for row in grid:
            for orange in row:
                if orange == 1:
                    return -1

        return ans

    def bsf(self, grid: List[List[int]]):
        rows = len(grid)
        cols = len(grid[0])
        dirs = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        fl_oranges = []
        for i, row in enumerate(grid):
            for j, orange in enumerate(row):
                if self.visited[i][j]:
                    continue
                if orange == 2:
                    fl_oranges.append((i, j))
                    self.visited[i][j] = True

        has_fl = False
        for i, j in fl_oranges:
            for dir in dirs:
                new_row = i + dir[0]
                new_col = j + dir[1]
                if (
                    new_row >= 0
                    and new_row < rows
                    and new_col >= 0
                    and new_col < cols
                    and grid[new_row][new_col] == 1
                ):
                    grid[new_row][new_col] = 2
                    has_fl = True

        return has_fl

    # 思路2: 用队列模拟
    def orangesRotting2(self, grid: List[List[int]]) -> int:
        visited = [[False for _ in row] for row in grid]
        dirs = [(-1, 0), (1, 0), (0, 1), (0, -1)]
        queue = collections.deque()
        for i, row in enumerate(grid):
            for j, orange in enumerate(row):
                if orange == 2:
                    queue.append((i, j))

        ans = 0
        rows = len(grid)
        cols = len(grid[0])
        while len(queue) > 0:
            size = len(queue)
            has_fl = False
            for _ in range(size):
                i, j = queue.popleft()
                if visited[i][j]:
                    continue
                for row_offset, col_offset in dirs:
                    new_row = i + row_offset
                    new_col = j + col_offset
                    if (
                        new_row >= 0
                        and new_row < rows
                        and new_col >= 0
                        and new_col < cols
                        and grid[new_row][new_col] == 1
                    ):
                        grid[new_row][new_col] = 2
                        has_fl = True
                        queue.append((new_row, new_col))
                visited[i][j] = True
            if has_fl:
                ans += 1

        for row in grid:
            for orange in row:
                if orange == 1:
                    return -1

        return ans
