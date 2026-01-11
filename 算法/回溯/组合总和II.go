package 回溯算法

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var path []int
	var used []bool
	used = make([]bool, len(candidates))
	sort.Ints(candidates)
	backTracking6(candidates, target, 0, 0, path, used, &ans)
	return ans
}

func backTracking6(candidates []int, target int, sum int, start int, path []int, used []bool, ans *[][]int) {
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	}
	for i := start; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			continue
		}
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		sum += candidates[i]
		path = append(path, candidates[i])
		used[i] = true
		backTracking6(candidates, target, sum, i+1, path, used, ans)
		sum -= candidates[i]
		path = path[:len(path)-1]
		used[i] = false
	}
}
