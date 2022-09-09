package main

import (
	"fmt"
)

func orderlyQueue(s string, k int) string {
	if len(s) == 1 {
		return s
	}
	news := maxStr(s, k)
	if news >= s {
		return s
	} else {
		return orderlyQueue(news, k)
	}
}

func maxStr(s string, num int) string {

	var k int
	var max, other string
	for i := 0; i < num; i++ {
		if max == "" {
			max = s[i : i+1]
			k = i
		} else {
			if max < s[i:i+1] {
				max = s[i : i+1]
				k = i
			}
		}
	}
	fmt.Printf("%s %d %d \n", s, num, k)
	if k == 0 {
		other = s[1:]
	} else {
		other = s[0:k] + s[k+1:]
	}
	return other + max
}

func main() {
	fmt.Println(orderlyQueue("kuh", 1))
	fmt.Println(orderlyQueue("baaca", 3))
}
