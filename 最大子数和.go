package main

/**
https://leetcode.cn/problems/maximum-subarray/

给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组是数组中的一个连续部分。

示例 1：

输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例 2：

输入：nums = [1]
输出：1
示例 3：

输入：nums = [5,4,-1,7,8]
输出：23


提示：

1 <= nums.length <= 105
-104 <= nums[i] <= 104


进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
*/

// 遍历所有和的情况，取最大和
// 复杂度O(N^3)
func maxSubArray1(nums []int) int {
	ans := -10001
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			if sum > ans {
				ans = sum
			}
		}
	}

	return ans
}

// 基于第一种方法剪枝，仍然是遍历
// 时间复杂度O(N^2)
func maxSubArray2(nums []int) int {
	ans := -10001
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
		}
		if sum > ans {
			ans = sum
		}
	}

	return ans
}

// dp[i] 表示在第i位的最大和
// 时间复杂度O(N)
func maxSubArray3(nums []int) int {
	// dp[i] = Max{nums[i], dp[i-1] + nums[i]}
	// dp[0] = nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	ans := -10001
	for i := 0; i < len(dp); i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}

// 分治，三种情况：1、左边取到最大和 2、右边取到最大和 3、横跨中间取到最大和
// 这类题都能用分治解决
// 时间复杂度O(NlogN)
func maxSubArray4(nums []int) int {
	return maxSumRec(nums, 0, len(nums)-1)
}

func maxSumRec(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) >> 1
	maxLeft := maxSumRec(nums, left, mid)
	maxRight := maxSumRec(nums, mid+1, right)

	// 横跨两边的情况
	maxLeftBorderSum, leftBorderSum := -10001, 0
	for i := mid; i >= left; i-- {
		leftBorderSum += nums[i]
		if leftBorderSum > maxLeftBorderSum {
			maxLeftBorderSum = leftBorderSum
		}
	}

	maxRightBorderSum, rightBorderSum := -10001, 0
	for i := mid + 1; i <= right; i++ {
		rightBorderSum += nums[i]
		if rightBorderSum > maxRightBorderSum {
			maxRightBorderSum = rightBorderSum
		}
	}

	return max3(maxLeft, maxRight, maxLeftBorderSum+maxRightBorderSum)
}

func max3(a int, b int, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}
