package main

import (
	"fmt"
)

// https://leetcode.cn/problems/beautiful-arrangement-ii/submissions/

// n=4  k=2  =>      [1,2,4,3]
// n=4  k=3  =>      [1,4,2,3]
// n=5  k=3  =>      [1,2,5,3,4]  1,3,2,1     处理节点是3
// n=5  k=4  =>      [1,5,2,4,3]   4,3,2,1    处理节点是2
// n=6  k=5   =>     [1,6,2,5,3,4] 5,4,3,2,1  处理节点是2
func constructArray(n int, k int) []int {
	var array []int
	// 计算处理节点
	begin := n - k + 1
	for i := 1; i <= n; i++ {
		// 如果i小于节点，直接加入数组
		if i < begin {
			array = append(array, i)
		} else {
			// 到达处理节点，按n->1递减和begin->n递增的顺序交替加入数组
			if (i-begin)%2 == 0 {
				array = append(array, n-(i-begin)/2)
			} else {
				array = append(array, begin+(i-begin)/2)
			}
		}
	}
	return array
}

// 更好的算法
func constructArrayPerfect(n, k int) []int {
	ans := make([]int, 0, n)
	for i := 1; i < n-k; i++ {
		ans = append(ans, i)
	}
	for i, j := n-k, n; i <= j; i++ {
		ans = append(ans, i)
		if i != j {
			ans = append(ans, j)
		}
		j--
	}
	return ans
}

func main() {
	fmt.Println(constructArray(3, 1))
	fmt.Println(constructArray(3, 2))
	fmt.Println(constructArray(5, 4))
}
