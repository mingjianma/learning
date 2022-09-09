package main

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/rearrange-spaces-between-words/

func reorderSpaces(text string) string {
	var wordList []string // 单词列表
	var str string        // 当前单词
	var space int         // 空格数量
	// 筛选单词和计算空格数
	for i := 0; i < len(text); i++ {
		if string(text[i]) == " " {
			space++
			if str != "" {
				str, wordList = "", append(wordList, str)
			}
		} else {
			str += string(text[i])
		}
	}
	// 检查最后是否是单词结尾
	if str != "" {
		str, wordList = "", append(wordList, str)
	}
	// 单词数量为1，空格置尾
	if len(wordList) == 1 {
		return wordList[0] + strings.Repeat(" ", space)
	}
	// 计算间隔空格数量和尾空格数
	split, tail := space/(len(wordList)-1), space%(len(wordList)-1)
	// 生成目标字符串
	for _, v := range wordList {
		if str == "" {
			str = v
		} else {
			str += strings.Repeat(" ", split) + v
		}
	}
	return str + strings.Repeat(" ", tail)
}

func main() {
	fmt.Println(reorderSpaces("  this   is  a sentence "))
	fmt.Println(reorderSpaces(" practice   makes   perfect"))
	fmt.Println(reorderSpaces("hello   world"))
}
