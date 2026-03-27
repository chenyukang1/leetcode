# https://leetcode.cn/problems/partition-labels/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def partitionLabels(self, s: str) -> List[int]:
        last = {}
        for i in range(len(s) - 1, -1, -1):
            if s[i] in last:
                continue
            last[s[i]] = i

        ans = []
        i = 0
        while i < len(s):
            max_reach = last[s[i]]

            j = i
            while j <= max_reach:
                if last[s[j]] > max_reach:
                    max_reach = last[s[j]]
                j += 1

            ans.append(max_reach - i + 1)
            i = max_reach + 1

        return ans


if __name__ == "__main__":
    s = Solution()
    print(s.partitionLabels("ababcbacadefegdehijhklij"))
