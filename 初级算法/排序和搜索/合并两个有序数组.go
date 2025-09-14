package 排序和搜索

// 归并
func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := 0, 0, 0
	tmp := make([]int, m)
	for idx := 0; idx < m; idx++ {
		tmp[idx] = nums1[idx]
	}
	for i < m && j < n {
		if tmp[i] <= nums2[j] {
			nums1[k] = tmp[i]
			k++
			i++
		} else {
			nums1[k] = nums2[j]
			k++
			j++
		}
	}
	for ; i < m; i, k = i+1, k+1 {
		nums1[k] = tmp[i]
	}
	for ; j < n; j, k = j+1, k+1 {
		nums1[k] = nums2[j]
	}
}

// 从后往前
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i > 0 && j > 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
		} else {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}
	for i > 0 {
		nums1[k] = nums1[i]
		i--
		k--
	}
	for j > 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
