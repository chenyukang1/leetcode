package 回溯算法

func combine(n int, k int) [][]int {
	var ans [][]int
	var path []int
	backTracking(n, k, 1, path, &ans)
	return ans
}

func backTracking(n int, k int, start int, path []int, ans *[][]int) {
	pathLen := len(path)
	if pathLen == k {
		tmp := make([]int, k)
		copy(tmp, path)
		*ans = append(*ans, tmp)
		return
	}

	for i := start; i <= n-pathLen+1; i++ {
		path = append(path, i)
		backTracking(n, k, i+1, path, ans)
		path = path[:len(path)-1]
	}
}
