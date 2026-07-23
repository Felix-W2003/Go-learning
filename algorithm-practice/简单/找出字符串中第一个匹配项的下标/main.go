package main

import (
	"fmt"
)

/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。
 示例 1：

输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
示例 2：

输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。


提示：

1 <= haystack.length, needle.length <= 104
haystack 和 needle 仅由小写英文字符组成
*/

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	// 目标串更长，直接返回-1
	if m > n {
		return -1
	}
	// i最大取值：n - m，避免i+j越界
	for i := 0; i <= n-m; i++ {
		match := true
		for j := 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break // 有一个不匹配，直接跳出内层
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func main() {
	str1 := "sadbutsad"
	str2 := "but"
	fmt.Println(strStr(str1, str2)) // 预期0

	str3 := "leetcode"
	str4 := "leeto"
	fmt.Println(strStr(str3, str4)) // 预期-1
}
