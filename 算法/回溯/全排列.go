package 回溯

func permute(nums []int) [][]int {
	var path []int
	var ans [][]int
	visited := make([]bool, len(nums))
	permuteRecurse(path, visited, nums, &ans)
	return ans
}

func permuteRecurse(path []int, visited []bool, nums []int, ans *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		path = append(path, nums[i])
		visited[i] = true
		permuteRecurse(path, visited, nums, ans)
		path = path[:len(path)-1]
		visited[i] = false
	}
}
