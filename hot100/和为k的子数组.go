package hot100

//https://leetcode.cn/problems/subarray-sum-equals-k/?envType=study-plan-v2&envId=top-100-liked

// 回溯，超时
func subarraySum(nums []int, k int) int {
	var ans int
	var path []int
	backTracking(0, nums, k, path, &ans)
	return ans
}

func backTracking(start int, nums []int, k int, path []int, ans *int) {
	if len(path) > 0 && sum(nums, path) == k {
		*ans++
	}
	for i := start; i < len(nums); i++ {
		if i > 0 && len(path) > 0 && path[len(path)-1] != i-1 {
			continue
		}
		path = append(path, i)
		backTracking(i+1, nums, k, path, ans)
		path = path[:len(path)-1]
	}
}

func sum(nums []int, path []int) int {
	ans := 0
	for _, i := range path {
		ans += nums[i]
	}
	return ans
}

// k = pre-(pre-k)
func subarraySum2(nums []int, k int) int {
	var ans int

	// init pre
	pre := make([]int, len(nums))
	count := make(map[int]int)
	count[0] = 1
	for i, num := range nums {
		if i == 0 {
			pre[i] = num
		} else {
			pre[i] = pre[i-1] + num
		}
		if count[pre[i]-k] > 0 {
			ans += count[pre[i]-k]
		}
		count[pre[i]]++
	}
	return ans
}
