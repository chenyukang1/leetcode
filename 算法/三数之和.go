package hot100

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/3sum/description
// 哈希表+去重
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		m := make(map[int]int)
		seen := make(map[string]struct{})
		for j := i + 1; j < len(nums); j++ {
			if k, ok := m[target-nums[j]]; ok {
				ints := []int{nums[i], nums[j], nums[k]}
				key := fmt.Sprint(ints)
				if _, ok = seen[key]; !ok {
					seen[key] = struct{}{}
					res = append(res, ints)
				}
			}
			m[nums[j]] = j
		}
	}
	return res
}

// 双指针
func threeSum2(nums []int) [][]int {
	sort.Ints(nums) // -4 -1 -1 0 1 2

	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r]
			if sum < target {
				l++
			} else if sum > target {
				r--
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			}
			for l < r && nums[l] == nums[l+1] {
				l++
			}
			for l < r && nums[r-1] == nums[r] {
				r--
			}
		}
	}
	return res
}
