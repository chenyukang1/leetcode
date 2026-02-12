from typing import List


class Solution:
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
                    fl_oranges.append((i,j))
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
