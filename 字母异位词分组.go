package main

import (
	"sort"
)

/**
https://leetcode.cn/problems/group-anagrams

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]

输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

解释：

在 strs 中没有字符串可以通过重新排列来形成 "bat"。
字符串 "nat" 和 "tan" 是字母异位词，因为它们可以重新排列以形成彼此。
字符串 "ate" ，"eat" 和 "tea" 是字母异位词，因为它们可以重新排列以形成彼此。
示例 2:

输入: strs = [""]

输出: [[""]]

示例 3:

输入: strs = ["a"]

输出: [["a"]]


提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母
*/

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		sorted := sortString(str)
		m[sorted] = append(m[sorted], str)
	}

	// 最终结果
	var ans [][]string
	for _, group := range m {
		ans = append(ans, group)
	}
	return ans
}

func sortString(s string) string {
	// 将字符串转换为 rune 切片，支持 Unicode 字符
	runes := []rune(s)

	// 按照字符排序
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// 拼接成字符串
	return string(runes)
}
