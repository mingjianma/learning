package main

import (
	"fmt"
	"strings"
)

//给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
//字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
func main() {
	s := "aaaaaa"
	t := "bbaaaa"
	fmt.Println(isSubsequence(s, t))
}

//匹配子串的每一字符是否在t中
//在t找到第一个字符以后从找到位置的后一个字符开始重新匹配
func isSubsequence(s string, t string) bool {
	result := true
	for _, v := range s {
		firstIndex := strings.Index(t, string(v))
		if firstIndex >= 0 {
			t = t[firstIndex+1:]
		} else {
			result = false
			break
		}
	}

	return result
}
