package main

import (
	"fmt"
)

func minOperations(logs []string) int {
	var num int
	// 遍历文件夹
	for _, v := range logs {
		// 当前目录操作跳过
		if v == "./" {
			continue
		} else if v == "../" {
			// 如果当前在根目录，上一层操作跳过，否则，层级减1
			if num == 0 {
				continue
			} else {
				num--
			}
		} else {
			// 其他操作层级加1
			num++
		}
	}
	return num
}

func main() {
	fmt.Println(minOperations([]string{"d1/", "d2/", "../", "d21/", "./"}))
	fmt.Println(minOperations([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}))
	fmt.Println(minOperations([]string{"d1/", "../", "../", "../"}))
	fmt.Println(minOperations([]string{"./", "wz4/", "../", "mj2/", "../", "../", "ik0/", "il7/"}))
}
