from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def __init__(self) -> None:
        self.mem = {}

    def rob(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        if not root.left and not root.right:
            return root.val
        if root in self.mem:
            return self.mem[root]

        val1 = root.val
        if root.left:
            val1 += self.rob(root.left.left) + self.rob(root.left.right)
        if root.right:
            val1 += self.rob(root.right.left) + self.rob(root.right.right)

        val2 = self.rob(root.left) + self.rob(root.right)
        res = max(val1, val2)
        self.mem[root] = res

        return res


if __name__ == "__main__":
    s = Solution()
    n = TreeNode(1)
    n.left = TreeNode(2)
    n.right = TreeNode(3)
    print(s.rob(n))
