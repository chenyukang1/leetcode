package graph

// https://leetcode.cn/problems/all-paths-from-source-to-target/description/

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	var path []int
	path = append(path, 0)
	allPathsSourceTargetDsf(0, path, &res, graph)
	return res
}

func allPathsSourceTargetDsf(start int, path []int, res *[][]int, graph [][]int) {
	if start == len(graph)-1 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for j := 0; j < len(graph[start]); j++ {
		path = append(path, graph[start][j])
		allPathsSourceTargetDsf(graph[start][j], path, res, graph)
		path = path[:len(path)-1]
	}
}
