package 数组

/**
给你两个整数数组nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]

提示：

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000

进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果nums1的大小比nums2 小，哪种方法更优？
如果nums2的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/

func intersect(nums1 []int, nums2 []int) []int {
	arr1 := make([]int, 1001)
	for _, num := range nums1 {
		arr1[num]++
	}

	arr2 := make([]int, 1001)
	for _, num := range nums2 {
		arr2[num]++
	}

	var ans []int
	for i := 0; i < 1001; i++ {
		m := min2(arr1[i], arr2[i])
		for m > 0 {
			ans = append(ans, i)
			m--
		}
	}
	return ans
}

func min2(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
