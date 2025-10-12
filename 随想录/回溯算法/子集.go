package 回溯算法

//https://leetcode.cn/problems/subsets/description/

func subsets(nums []int) [][]int {
	var ans [][]int
	var path []int
	ans = append(ans, []int{})
	backTracking9(0, nums, path, &ans)
	return ans
}

func backTracking9(start int, nums []int, path []int, ans *[][]int) {
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		backTracking9(i+1, nums, path, ans)
		path = path[:len(path)-1]
	}
}
