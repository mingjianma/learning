package main

import (
	"fmt"
	"sort"
)

/*
	698. 划分为k个相等的子集
*/

// 回溯
// index : 第 index 个球开始做选择
// bucket : 桶
func backtrack(nums []int, index int, bucket []int, k int, target int) bool {
	// 递减排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] >= nums[j]
	})
	// 1.结束条件
	if index == len(nums) {
		// 判断现在桶中的球是否符合要求 -> 路径是否满足要求
		for i := 0; i < k; i++ {
			if bucket[i] != target {
				return false
			}
			return true
		}
	}
	// 2.nums[index]球开始做选择
	for i := 0; i < k; i++ {
		// 剪枝：放入球后超过 target 的值，选择一下桶
		if bucket[i]+nums[index] > target {
			continue
		}
		// 如果当前桶和上一个桶内的元素和相等，则跳过
		// 原因：如果元素和相等，那么 nums[index] 选择上一个桶和选择当前桶可以得到的结果是一致的
		if i > 0 && bucket[i] == bucket[i-1] {
			continue
		}
		// 对于第一个球，任意放到某个桶中的效果都是一样的，所以我们规定放到第一个桶中
		if i > 0 && index == 0 {
			break
		}
		// 做选择：放入 i 号桶
		bucket[i] += nums[index]
		// 处理下一个球，即 nums[index + 1],回溯
		if backtrack(nums, index+1, bucket, k, target) {
			return true
		}
		// 撤销上一操作
		bucket[i] -= nums[index]
	}
	// 3.k 个桶都不满足要求
	return false
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	// 目标数
	target := sum / k
	// 桶个数
	bucket := make([]int, k+1)
	return backtrack(nums, 0, bucket, k, target)
}

func main() {
	// fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	// fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
	fmt.Println(canPartitionKSubsets([]int{1, 1, 1, 1, 2, 2, 2, 2}, 4))
}
