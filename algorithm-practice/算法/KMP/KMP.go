package main

import "fmt"

//KMPSearch kmp查找，返回模式串pattern在text中首次出现下标，不存在返回-1
func KMPSearch(text, pattern string) int {
	n := len(text)
	m := len(pattern)
	if m == 0 {
		return 0
	}
	if n < m {
		return -1
	}

	next := buildNext(pattern)
	i := 0 //text指针
	j := 0 //pattern指针
	for i < n {
		if text[i] == pattern[j] {
			i++
			j++
			if j == m {
				//匹配成功，返回起始位置
				return i - j
			}
		} else {
			if j != 0 {
				j = next[j-1]

			} else {
				i++
			}
		}
	}
	return -1
}

// buildNext 构建next数组（标准前缀函数）
func buildNext(pattern string) []int {
	m := len(pattern)
	next := make([]int, m)
	length := 0 //当前最长相等前后缀长度
	i := 1
	for i < m {
		if pattern[i] == pattern[length] {
			length++
			next[i] = length
			i++
		} else {
			if length != 0 {
				length = next[length-1]
			} else {
				next[i] = 0
				i++
			}
		}
	}
	return next
}

func main() {
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	pos := KMPSearch(text, pattern)
	if pos != -1 {
		fmt.Printf("匹配成功，起始索引：%d\n", pos)
		fmt.Printf("截取匹配内容：%s\n", text[pos:pos+len(pattern)])
	} else {
		fmt.Println("未找到匹配字符串")
	}

	// 测试案例2
	fmt.Println(KMPSearch("aaaaa", "bba"))
	fmt.Println(KMPSearch("abcabcabc", "abc"))
}
