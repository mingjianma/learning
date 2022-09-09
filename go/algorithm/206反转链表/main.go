package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var reverseList *ListNode
	for head != nil {
		if reverseList == nil {
			reverseList = &ListNode{
				Val:  head.Val,
				Next: nil,
			}
		} else {
			newNode := &ListNode{
				Val:  head.Val,
				Next: reverseList,
			}
			reverseList = newNode
		}
		head = head.Next
	}

	return reverseList
}

func main() {
	val := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	result := reverseList(val)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
