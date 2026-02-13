# https://leetcode.cn/problems/course-schedule/description/?envType=study-plan-v2&envId=top-100-liked

import collections
from typing import List


# 思路1: 计算入度，入度为0则可以先修，循环修完所有可以先修课程，最后比较是否能够修完
# 思路2: 计算链表成环，做不到，因为一门课可能有多个先修课，链表无法解决
class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        adj = [[] for _ in range(numCourses)]
        indegree = [0] * numCourses

        for cur, pre in prerequisites:
            adj[pre].append(cur)
            indegree[cur] += 1

        queue = collections.deque([i for i in range(numCourses) if indegree[i] == 0])

        finished = 0
        while len(queue) > 0:
            pre = queue.popleft()
            finished += 1
            for cur in adj[pre]:
                indegree[cur] -= 1
                if indegree[cur] == 0:
                    queue.append(cur)

        return finished >= numCourses
