package 回溯算法

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var path []int
	backTracking5(candidates, target, 0, 0, path, &ans)
	return ans
}

func backTracking5(candidates []int, target int, sum, start int, path []int, ans *[][]int) {
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
	}
	for i := start; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			continue
		}
		sum += candidates[i]
		path = append(path, candidates[i])
		backTracking5(candidates, target, sum, i, path, ans)
		sum -= candidates[i]
		path = path[:len(path)-1]
	}
}
