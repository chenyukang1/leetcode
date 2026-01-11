package 数组

func rotate(nums []int, k int) {
	k = k % len(nums)
	j := len(nums) - k - 1
	var arr []int

	for i, m := j, 0; i < len(nums); i++ {
		arr[m] = nums[i]
		m++
	}

	for i, m := j, 0; i < len(nums); i++ {
		nums[i] = nums[m]
		m++
	}

	for i := 0; i < j; i++ {
		nums[i] = arr[i]
	}
}

func rotate1(nums []int, k int) {
	k = k % len(nums)
	reserve(nums, 0, len(nums)-k-1)
	reserve(nums, len(nums)-k, len(nums)-1)
	reserve(nums, 0, len(nums)-1)
}

func reserve(nums []int, i int, j int) {
	for i < j {
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
		i++
		j--
	}
}
