package main

import "fmt"

/*给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
示例 1：

输入：s = "()"

输出：true

示例 2：

输入：s = "()[]{}"

输出：true

示例 3：

输入：s = "(]"

输出：false

示例 4：

输入：s = "([])"

输出：true

示例 5：

输入：s = "([)]"

输出：false

 提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

func isValid(s string) bool {
	//有括号映射对应的左括号
	match := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0)

	for _, char := range s {
		//左括号:入栈
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			//右括号:栈为空直接不合法
			if len(stack) == 0 {
				return false
			}
			//取出栈顶
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			//不匹配直接返回false
			if match[char] != top {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	testCases := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"((",
		"}",
		")))(((",
	}
	for _, str := range testCases {
		fmt.Printf("字符串：%s \n 是否有效：%t\n", str, isValid(str))
	}
}
