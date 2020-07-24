package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2, 3}
	result := intersect(nums1, nums2)
	fmt.Println(result)

	arr1 := []int{1, 2, 3, 4, 4, 13}
	arr2 := []int{1, 2, 3, 9, 10}
	result = intersectSort2(arr1, arr2)
	fmt.Println(result)
}

//两个无序数组的交集元素，同时应与两个数组中出现的次数一致。这样就导致了我们需要知道每个值出现的次数，所以映射关系就成了<元素,出现次数>
func intersect(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	for _, v := range nums1 {
		//遍历nums1，初始化map
		m0[v] += 1
	}
	k := 0
	for _, v := range nums2 {
		//如果元素相同，将其存入nums2中，并将出现次数减1
		if m0[v] > 0 {
			m0[v] -= 1
			nums2[k] = v
			k++
		}
	}
	return nums2[0:k]
}

//已排好序的数组交集 自己想的
func intersectSort1(nums1 []int, nums2 []int) []int {
	len1 := len(nums1)
	len2 := len(nums2)
	var p1, p2, k int
	for i := p1; i < len1; i++ {
		for j := p2; j < len2; j++ {
			if nums1[i] == nums2[j] {
				nums2[k] = nums1[i]
				p1 = i + 1
				p2 = j + 1
				k++
				break
			}
		}
		if p1 >= len1 && p2 >= len2 {
			break
		}
	}
	return nums2[0:k]
}

//已排好序的数组交集 更好的
func intersectSort2(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}
