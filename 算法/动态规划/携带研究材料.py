# https://kamacoder.com/problempage.php?pid=1046

# 01背包问题:二维数组
def bag():
    m, n = map(int, input().split())
    weight = [int(item) for item in input().split()]
    value = [int(item) for item in input().split()]
    # dp[i][j] 表示到i物品为止放入j空间，最大的价值
    dp = [[0 for _ in range(n + 1)] for _ in range(m + 1)]
    for i in range(m + 1):
        dp[0][i] = 0
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if j - weight[i - 1] >= 0:
                dp[i][j] = max(
                    dp[i - 1][j], dp[i - 1][j - weight[i - 1]] + value[i - 1]
                )
            else:
                dp[i][j] = dp[i - 1][j]

    return dp[m][n]


# 滚动数组
def bag2():
    m, n = map(int, input().split())
    weight = [int(item) for item in input().split()]
    value = [int(item) for item in input().split()]
    # dp[j] 表示容量为j的背包，所背物品的最大价值
    dp = [0 for _ in range(n + 1)]
    for i in range(0, m):
        for j in range(n, weight[i] - 1, -1):
            dp[j] = max(dp[j], dp[j - weight[i]] + value[i])
    return dp[n]
