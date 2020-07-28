package main

import "fmt"

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(Sum(nums, target))
}

//遍历nums，将每个Num的差值保存到map，检查每个num是否在map中，存在时返回对应的下标
func Sum(nums []int, target int) []int {
	resultMap := make(map[int]int)
	for k, v := range nums {
		if firstIndex, ok := resultMap[v]; ok {
			return []int{firstIndex, k}
		} else {
			resultMap[target-v] = k
		}
	}
	return []int{}
}
