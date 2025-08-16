package _62

//Q2.排序排列
//
//给你一个长度为 n 的整数数组 nums，其中 nums 是范围 [0..n - 1] 内所有数字的一个排列。
//
//你可以在满足条件 nums[i] AND nums[j] == k 的情况下交换下标i 和 j 的元素，其中 AND 表示按位与操作，k 是一个非负整数。
//
//返回可以使数组按非递减顺序排序的最大值 k（允许进行任意次这样的交换）。如果 nums 已经是有序的，返回 0。
//
//排列是数组所有元素的一种重新排列。

func sortPermutation(nums []int) int {
	// 判断是否已排序
	sorted := true
	for i, v := range nums {
		if v != i {
			sorted = false
			break
		}
	}
	if sorted {
		return 0
	}

	// 初始化 k 为全 1，根据最大位数决定，假设用 32 位
	k := ^0 // 注意：在 Go 中 ^0 是全 1 的整数（依赖于 int 位宽）
	for i, v := range nums {
		if v != i {
			k &= v
		}
	}
	return k
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
