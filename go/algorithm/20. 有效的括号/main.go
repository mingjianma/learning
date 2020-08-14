package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//有效字符串需满足：
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。

func main() {
	s := "()"
	fmt.Println(isValid(s))
	s = "()[]{}"
	fmt.Println(isValid(s))
	s = "(]"
	fmt.Println(isValid(s))
	s = "([)]"
	fmt.Println(isValid(s))
	s = "{[]}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	strMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var left []byte
	for i := 0; i < len(s); i++ {
		if val, ok := strMap[s[i]]; ok {
			left = append(left, val)
		} else {
			if len(left) == 0 {
				return false
			} else {
				if left[len(left)-1] != s[i] {
					return false
				}

				left = left[:len(left)-1]
			}
		}
	}
	if len(left) > 0 {
		return false
	} else {
		return true
	}
}
