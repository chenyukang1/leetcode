package 回溯算法

//https://leetcode.cn/problems/restore-ip-addresses/description/

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var ans []string
	var nums []string
	backTracking8(0, s, nums, &ans)
	return ans
}

// 取4个点，最后一个一定要在最后
func backTracking8(start int, s string, nums []string, ans *[]string) {
	if len(nums) == 4 && start == len(s) {
		*ans = append(*ans, strings.Join(nums, "."))
		return
	}
	for i := start + 1; i <= len(s); i++ {
		if i-start > 3 {
			break
		}
		if len(s[start:i]) > 1 && s[start:i][0] == '0' {
			break
		}
		num, err := strconv.Atoi(s[start:i])
		if err != nil || num < 0 || num > 255 {
			continue
		}
		nums = append(nums, s[start:i])
		backTracking8(i, s, nums, ans)
		nums = nums[:len(nums)-1]
	}
}

// 取3个点，最后一个再判断
func backTracking8_2(start int, s string, nums []string, ans *[]string) {
	if len(nums) == 3 {
		if len(s)-start > 3 {
			return
		}
		if len(s[start:]) > 1 && s[start:][0] == '0' {
			return
		}
		num, err := strconv.Atoi(s[start:])
		if err != nil || num < 0 || num > 255 {
			return
		}
		nums = append(nums, s[start:])
		*ans = append(*ans, strings.Join(nums, "."))
		return
	}
	for i := start + 1; i <= len(s); i++ {
		if i-start > 3 {
			break
		}
		if len(s[start:i]) > 1 && s[start:i][0] == '0' {
			break
		}
		num, err := strconv.Atoi(s[start:i])
		if err != nil || num < 0 || num > 255 {
			continue
		}
		nums = append(nums, s[start:i])
		backTracking8_2(i, s, nums, ans)
		nums = nums[:len(nums)-1]
	}
}
