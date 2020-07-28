package main

import "fmt"

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
func main() {
	/*str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
	str = "bbbbbb"
	fmt.Println(lengthOfLongestSubstring(str))*/
	str := "dvdf"
	fmt.Println(lengthOfLongestSubstring(str))
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	var p, maxLen int
	for maxLen < len(s) || s != "" {
		strMap := make(map[string]int)
		strLen := len(s)
		for i := p; i < strLen; i++ {
			if index, ok := strMap[string(s[i])]; ok {
				s = s[index+1:]
				p = 0
				if i-p > maxLen {
					maxLen = i - p
				}
				break
			} else {
				if strLen-1 == i {
					s = ""
					if i-p+1 > maxLen {
						maxLen = i - p + 1
					}
					break
				} else {
					strMap[string(s[i])] = i
				}
			}
		}
	}
	return maxLen
}

func lengthOfLongestSubstringNew(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			rk++
			m[s[rk]]++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
