package 回溯算法

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	var path []int
	used := make([]bool, len(nums))
	backTracking10(0, nums, path, used, &ans)
	return ans
}

func backTracking10(start int, nums []int, path []int, used []bool, ans *[][]int) {
	temp := make([]int, len(path))
	copy(temp, path)
	*ans = append(*ans, temp)
	for i := start; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backTracking10(i+1, nums, path, used, ans)
		path = path[:len(path)-1]
		used[i] = false
	}
}
