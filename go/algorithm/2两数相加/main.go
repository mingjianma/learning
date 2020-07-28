package main

import (
	"encoding/json"
	"fmt"
)

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val:  5,
		Next: nil,
	}

	l2 := &ListNode{
		Val:  5,
		Next: nil,
	}

	res := addTwoNumbers(l1, l2)
	jstr, _ := json.Marshal(res)
	fmt.Println(string(jstr))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//递归
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	var sum int
	//补完空next
	if l1 == nil {
		l1 = &ListNode{
			Val:  0,
			Next: nil,
		}
	} else if l2 == nil {
		l2 = &ListNode{
			Val:  0,
			Next: nil,
		}
	}

	sum = l1.Val + l2.Val
	//sum大于10时给next加1
	if sum >= 10 {
		sum -= 10
		if l1.Next != nil {
			l1.Next.Val += 1
		} else if l2.Next != nil {
			l2.Next.Val += 1
		} else {
			l1.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
		}
	}

	return &ListNode{
		Val:  sum,
		Next: addTwoNumbers(l1.Next, l2.Next),
	}
}

//循环
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		curr.Next = new(ListNode)
		curr = curr.Next
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		curr.Val = carry % 10
		carry /= 10
	}
	return dummy.Next
}
