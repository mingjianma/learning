package main

import (
	"fmt"
	"sort"
)

func specialArray(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i <= nums[j] {
				if i == len(nums)-j {
					return i
				}
				break
			}
		}
	}
	return -1
}

// 方法一：降序排序 + 一次遍历
// 思路与算法
//
// 我们可以首先将数组进行降序排序，这样一来，我们就可以通过遍历的方式得到数组的特征值了。
// 根据特征值 x 的定义，xx 一定是在 [1,n] 范围内的一个整数，其中 n 是数组 nums 的长度。
// 因此，我们可以遍历 [1,n] 并判断某个整数 ii 是否为特征值。
// 若 i 为特征值，那么nums 中恰好有 i 个元素大于等于 i。
// 由于数组已经降序排序，说明 nums[i−1] 必须大于等于 i，并且nums[i]（如果存在）必须小于 i。
func specialArrayPerfect(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, n := 1, len(nums); i <= n; i++ {
		if nums[i-1] >= i && (i == n || nums[i] < i) {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(specialArray([]int{0, 4, 3, 0, 4}))
	fmt.Println(specialArray([]int{0, 0}))
	fmt.Println(specialArray([]int{3, 5}))
	fmt.Println(specialArray([]int{3, 6, 7, 7, 0}))
}
