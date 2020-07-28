package main

import (
	"fmt"
	"strings"
)

//编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
//所有输入只包含小写字母 a-z
func main() {
	slice := []string{"flower", "flow", "flight"}
	str := longestPrefix(slice)
	fmt.Println(str)
	str1 := longestCommonPrefix(slice)
	fmt.Println(str1)
}

func longestPrefix(slice []string) string {
	if len(slice) < 1 {
		return ""
	}
	var result int
	first := len(slice[0])
	for i := 0; i <= first; i++ {
		var p string
		flag := true
		for _, v := range slice {
			if i > len(v) {
				flag = false
				break
			}

			if p == "" {
				p = string(v[i])
			} else {
				if p != string(v[i]) {
					flag = false
					break
				}
			}
		}
		if flag == false {
			break
		}
		result++
	}
	return slice[0][0:result]

}

//我们要想寻找最长公共前缀，那么首先这个前缀是公共的，我们可以从任意一个元素中找到它。假定我们现在就从一个数组中寻找最长公共前缀，
//那么首先，我们可以将第一个元素设置为基准元素x0。假如数组为["flow","flower","flight"]，flow就是我们的基准元素x0。
//然后我们只需要依次将基准元素和后面的元素进行比较（假定后面的元素依次为x1,x2,x3....），不断更新基准元素，直到基准元素和所有元素都满足最长公共前缀的条件，就可以得到最长公共前缀。
//具体比对过程如下：
//如果strings.Index(x1,x) == 0，则直接跳过（因为此时x就是x1的最长公共前缀），对比下一个元素。（如flower和flow进行比较）
//如果strings.Index(x1,x) != 0, 则截取掉基准元素x的最后一个元素，再次和x1进行比较，直至满足string.Index(x1,x) == 0，此时截取后的x为x和x1的最长公共前缀。
//（如flight和flow进行比较，依次截取出flow-flo-fl，直到fl被截取出，此时fl为flight和flow的最长公共前缀）
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, k := range strs {
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
