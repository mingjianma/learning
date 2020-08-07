package main

import "fmt"

//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
//请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//你可以假设 nums1 和 nums2 不会同时为空。
func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Print(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Print(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{3}
	nums2 = []int{-2, -1}
	fmt.Print(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len3 := len(nums1) + len(nums2)
	nums3 := make([]int, len3)
	for k := range nums3 {
		len1 := len(nums1)
		len2 := len(nums2)
		if len1 == 0 {
			nums3[k] = nums2[0]
			nums2 = nums2[1:]
			continue
		}

		if len2 == 0 {
			nums3[k] = nums1[0]
			nums1 = nums1[1:]
			continue
		}

		val1 := nums1[0]
		val2 := nums2[0]
		if val1 < val2 {
			nums3[k] = val1
			if len1 > 1 {
				nums1 = nums1[1:]
			} else {
				nums1 = nil
			}
		} else {
			nums3[k] = val2
			if len2 > 1 {
				nums2 = nums2[1:]
			} else {
				nums2 = nil
			}
		}
	}

	if (len3)%2 == 1 {
		return float64(nums3[(len3+1)/2-1])
	} else {
		return (float64(nums3[(len3)/2-1]) + float64(nums3[(len3)/2])) / 2
	}
}
