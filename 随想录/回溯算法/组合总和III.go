package 回溯算法

func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var path []int
	backTracking3(n, k, 1, path, &ans)
	return ans
}

func backTracking3(n int, k int, start int, path []int, ans *[][]int) {
	if len(path) == k {
		if sum(path) == n {
			*ans = append(*ans, path)
			return
		}
	}

	for i := start; i <= 9; i++ {
		if sum(path) > n {
			return
		}
		path = append(path, i)
		backTracking3(n, k, i+1, path, ans)
		path = path[:len(path)-1]
	}
}

func sum(arr []int) int {
	s := 0
	for _, i := range arr {
		s += i
	}
	return s
}
