package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	// 将数字转换字符串
	numStr := []byte(strconv.Itoa(num))
	for j := 0; j < len(numStr)-1; j++ {
		var max = j
		// 检测是否有比当前位置大的数字
		for i := len(numStr) - 1; i > j; i-- {
			if numStr[i] > numStr[max] {
				max = i
			}
		}
		// 没有检查下一个字符
		if numStr[max] == numStr[j] {
			continue
		}
		// 有的话交换两个数字的位置
		numStr[j], numStr[max] = numStr[max], numStr[j]
		break
	}
	// 将字符串转回数字
	newNum, _ := strconv.Atoi(string(numStr))
	return newNum
}

func main() {
	fmt.Println(maximumSwap(2736))
	fmt.Println(maximumSwap(9973))
	fmt.Println(maximumSwap(98368))
}
